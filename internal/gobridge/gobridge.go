// Package gobridge understands the surface of a 'use go' block: which
// functions it declares, how many parameters each takes, what it returns, and
// whether it can fail (returns (T, error)). Both the Watcher and the
// transpiler use this so Go-backed actions behave exactly like native ones.
package gobridge

import (
	"regexp"
	"strings"
)

// Func describes one Go function declared in a use-go block.
type Func struct {
	Name     string // the Go name, e.g. "luckyNumber" (callable as "lucky number")
	Params   int
	Ret      string // "int", "float64", "string", "bool", "" for none/unknown
	Fallible bool   // returns (T, error) or just error
}

var funcRe = regexp.MustCompile(`(?m)^\s*func\s+([A-Za-z_][A-Za-z0-9_]*)\s*\(([^)]*)\)\s*(?:\(([^)]*)\)|([\w\[\]\.\*]+))?`)

// Imports pulls single-line `import "pkg"` declarations out of a block and
// returns them plus the block with those lines removed.
func Imports(code string) (imports []string, rest string) {
	var kept []string
	for _, line := range strings.Split(code, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "import \"") && strings.HasSuffix(t, "\"") {
			imports = append(imports, strings.Trim(strings.TrimPrefix(t, "import"), " \""))
			continue
		}
		kept = append(kept, line)
	}
	return imports, strings.Join(kept, "\n")
}

// Funcs parses every top-level func declaration in a block.
func Funcs(code string) []Func {
	var out []Func
	for _, m := range funcRe.FindAllStringSubmatch(code, -1) {
		f := Func{Name: m[1], Params: countParams(m[2])}
		rets := m[3]
		if rets == "" {
			rets = m[4]
		}
		parts := splitTop(rets)
		for i, p := range parts {
			p = strings.TrimSpace(p)
			// a named return like "n int" -> take the last word as the type
			if ws := strings.Fields(p); len(ws) > 1 {
				p = ws[len(ws)-1]
			}
			if p == "error" {
				f.Fallible = true
			} else if i == 0 {
				f.Ret = p
			}
		}
		out = append(out, f)
	}
	return out
}

func countParams(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	return len(splitTop(s))
}

// splitTop splits on commas (no nesting concerns for the simple signatures we
// support in v0.4 — documented in the library guide).
func splitTop(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
