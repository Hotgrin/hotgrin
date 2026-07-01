// Package loader turns a SimpleScript source file plus everything it `use`s
// into a single combined program, ready for the checker and transpiler.
//
// A library is just another .ss file. When a program uses one, the library's
// actions are merged into the program (whole-program transpile). Paths are
// resolved relative to the importing file, and each file is loaded once even if
// several files use it. Remote (e.g. GitHub) libraries are not fetched yet.
package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hotgrin/simplescript/internal/ast"
	"github.com/hotgrin/simplescript/internal/lexer"
	"github.com/hotgrin/simplescript/internal/parser"
)

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
		} else if _, ok := s.(*ast.ActionStmt); ok {
			// A library contributes its actions to the whole program.
			l.libActions = append(l.libActions, s)
		}
		// Non-action top-level statements in a library are ignored.
	}
	return out
}

func (l *loader) loadUse(u *ast.UseStmt, fromDir string) {
	if strings.Contains(u.Path, "://") || strings.HasPrefix(u.Path, "github.com/") {
		l.errs = append(l.errs,
			fmt.Sprintf("remote libraries aren't supported yet (%q) — use a local .ss file path", u.Path))
		return
	}
	p := u.Path
	if !strings.HasSuffix(p, ".ss") {
		p += ".ss"
	}
	if !filepath.IsAbs(p) {
		p = filepath.Join(fromDir, p)
	}
	l.load(p, false)
}
