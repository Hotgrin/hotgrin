// Package loader turns a hotgrin source file plus everything it `use`s
// into a single combined program, ready for the checker and transpiler.
//
// A library is just another .hot file. When a program uses one, the library's
// actions are merged into the program (whole-program transpile). Paths are
// resolved relative to the importing file, and each file is loaded once even if
// several files use it. Remote (e.g. GitHub) libraries are not fetched yet.
package loader

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hotgrin/hotgrin/internal/ast"
	"github.com/hotgrin/hotgrin/internal/lexer"
	"github.com/hotgrin/hotgrin/internal/parser"
)

//go:embed std/*.hot
var stdFS embed.FS

type loader struct {
	visited    map[string]bool
	libActions []ast.Stmt
	errs       []string
}

// LoadFile loads mainPath and everything it uses, returning one combined
// program (library actions first, then the main file's own statements) and any
// problems found along the way.
func LoadFile(mainPath string) (*ast.Program, []string) {
	l := &loader{visited: map[string]bool{}}
	mainStmts := l.load(mainPath, true)
	combined := append([]ast.Stmt{}, l.libActions...)
	combined = append(combined, mainStmts...)
	return &ast.Program{Statements: combined}, l.errs
}

func (l *loader) load(path string, isMain bool) []ast.Stmt {
	abs, err := filepath.Abs(path)
	if err != nil {
		abs = path
	}
	if l.visited[abs] {
		return nil // already loaded (handles cycles and diamonds)
	}
	l.visited[abs] = true

	data, err := os.ReadFile(path)
	if err != nil {
		l.errs = append(l.errs, fmt.Sprintf("could not open library %q", path))
		return nil
	}

	tokens := lexer.New(string(data)).Tokenize()
	prog, perrs := parser.New(tokens).Parse()
	for _, e := range perrs {
		l.errs = append(l.errs, fmt.Sprintf("%s line %d: %s", filepath.Base(path), e.Line, e.Message))
	}

	dir := filepath.Dir(path)
	var out []ast.Stmt
	for _, s := range prog.Statements {
		if u, ok := s.(*ast.UseStmt); ok {
			l.loadUse(u, dir)
			continue
		}
		if isMain {
			out = append(out, s)
		} else {
			switch s.(type) {
			case *ast.ActionStmt, *ast.GoBlockStmt:
				// A library contributes its actions (and go blocks).
				l.libActions = append(l.libActions, s)
			}
		}
		// Non-action top-level statements in a library are ignored.
	}
	return out
}

// loadSource parses library source that didn't come from a local file (the
// embedded standard library) and merges its actions and go-blocks.
func (l *loader) loadSource(src, label string) {
	tokens := lexer.New(src).Tokenize()
	prog, perrs := parser.New(tokens).Parse()
	for _, e := range perrs {
		l.errs = append(l.errs, fmt.Sprintf("%s line %d: %s", label, e.Line, e.Message))
	}
	for _, st := range prog.Statements {
		switch st.(type) {
		case *ast.ActionStmt, *ast.GoBlockStmt:
			l.libActions = append(l.libActions, st)
		}
	}
}

// loadRemote fetches a library from GitHub with git (cached under the user's
// home directory) and loads its entry file.
func (l *loader) loadRemote(spec string) {
	repoSpec, tag := spec, ""
	if at := strings.LastIndex(spec, "@"); at > 0 {
		repoSpec, tag = spec[:at], spec[at+1:]
	}
	segs := strings.Split(repoSpec, "/")
	if len(segs) < 3 {
		l.errs = append(l.errs, fmt.Sprintf("a remote library needs the form github.com/user/repo, got %q", spec))
		return
	}
	repo := strings.Join(segs[:3], "/")
	sub := strings.Join(segs[3:], "/")

	home, err := os.UserHomeDir()
	if err != nil {
		home = os.TempDir()
	}
	cache := filepath.Join(home, ".hotgrin", "cache", filepath.FromSlash(repo))
	if tag != "" {
		cache += "@" + tag
	}
	if _, err := os.Stat(cache); err != nil {
		if _, err := exec.LookPath("git"); err != nil {
			l.errs = append(l.errs, "fetching remote libraries needs git installed — get it from https://git-scm.com")
			return
		}
		args := []string{"clone", "--quiet", "--depth", "1"}
		if tag != "" {
			args = append(args, "--branch", tag)
		}
		args = append(args, "https://"+repo+".git", cache)
		if out, err := exec.Command("git", args...).CombinedOutput(); err != nil {
			l.errs = append(l.errs, fmt.Sprintf("could not fetch %q: %s", repo, strings.TrimSpace(string(out))))
			return
		}
	}

	entry := filepath.Join(cache, filepath.FromSlash(sub))
	if sub == "" {
		entry = filepath.Join(cache, "lib.hot")
	} else if fi, err := os.Stat(entry); err == nil && fi.IsDir() {
		entry = filepath.Join(entry, "lib.hot")
	} else if !strings.HasSuffix(entry, ".hot") {
		entry += ".hot"
	}
	if _, err := os.Stat(entry); err != nil {
		l.errs = append(l.errs, fmt.Sprintf("library %q has no %s", spec, filepath.Base(entry)))
		return
	}
	l.load(entry, false)
}

func (l *loader) loadUse(u *ast.UseStmt, fromDir string) {
	// Standard libraries ship inside hotgrin itself: use "std/text"
	if strings.HasPrefix(u.Path, "std/") {
		name := strings.TrimSuffix(strings.TrimPrefix(u.Path, "std/"), ".hot")
		key := "std:" + name
		if l.visited[key] {
			return
		}
		l.visited[key] = true
		data, err := stdFS.ReadFile("std/" + name + ".hot")
		if err != nil {
			l.errs = append(l.errs, fmt.Sprintf("there is no standard library called %q", u.Path))
			return
		}
		l.loadSource(string(data), "std/"+name+".hot")
		return
	}
	// Remote libraries: use tools from "github.com/user/repo[/sub/path][@tag]"
	if strings.HasPrefix(u.Path, "github.com/") {
		l.loadRemote(u.Path)
		return
	}
	if strings.Contains(u.Path, "://") {
		l.errs = append(l.errs,
			fmt.Sprintf("only github.com/... remote libraries are supported for now (%q)", u.Path))
		return
	}
	p := u.Path
	if !strings.HasSuffix(p, ".hot") {
		p += ".hot"
	}
	if !filepath.IsAbs(p) {
		p = filepath.Join(fromDir, p)
	}
	l.load(p, false)
}
