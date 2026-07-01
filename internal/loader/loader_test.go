package loader

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hotgrin/simplescript/internal/ast"
)

func write(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func actionNames(p *ast.Program) []string {
	var names []string
	for _, s := range p.Statements {
		if a, ok := s.(*ast.ActionStmt); ok {
			names = append(names, a.Name)
		}
	}
	return names
}

func TestLoadMergesLibraryActions(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "lib.ss", "action greet with who\ngive back who\nend action")
	write(t, dir, "app.ss", "use \"lib\"\nsay greet with \"AJ\"")

	prog, errs := LoadFile(filepath.Join(dir, "app.ss"))
	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if got := actionNames(prog); len(got) != 1 || got[0] != "greet" {
		t.Errorf("expected library action 'greet' merged in, got %v", got)
	}
	for _, s := range prog.Statements {
		if _, ok := s.(*ast.UseStmt); ok {
			t.Error("use statement leaked into the combined program")
		}
	}
}

func TestTransitiveAndLoadOnce(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "base.ss", "action a\ngive back 1\nend action")
	write(t, dir, "mid.ss", "use \"base\"\naction b\ngive back a\nend action")
	write(t, dir, "app.ss", "use \"mid\"\nuse \"base\"\nsay b")

	prog, errs := LoadFile(filepath.Join(dir, "app.ss"))
	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	count := map[string]int{}
	for _, n := range actionNames(prog) {
		count[n]++
	}
	if count["a"] != 1 {
		t.Errorf("action 'a' should be loaded exactly once, got %d", count["a"])
	}
}

func TestRemoteLibraryRejected(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "app.ss", "use math from \"github.com/x/y\"\nsay 1")
	_, errs := LoadFile(filepath.Join(dir, "app.ss"))
	if len(errs) == 0 || !strings.Contains(strings.Join(errs, " "), "remote libraries") {
		t.Errorf("expected a remote-library error, got %v", errs)
	}
}

func TestMissingLibraryReported(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "app.ss", "use \"nope\"\nsay 1")
	_, errs := LoadFile(filepath.Join(dir, "app.ss"))
	if len(errs) == 0 {
		t.Error("expected an error for a missing library")
	}
}
