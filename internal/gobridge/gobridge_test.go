package gobridge

import "testing"

func TestImports(t *testing.T) {
	imps, rest := Imports("import \"strings\"\nfunc f() {}\nimport \"os\"")
	if len(imps) != 2 || imps[0] != "strings" || imps[1] != "os" {
		t.Errorf("imports = %v", imps)
	}
	if rest != "\nfunc f() {}\n" && rest != "func f() {}" {
		// exact whitespace shape may vary; the func must remain, imports must not
		if !contains(rest, "func f()") || contains(rest, "import") {
			t.Errorf("rest = %q", rest)
		}
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})())
}

func TestFuncs(t *testing.T) {
	fs := Funcs(`func luckyNumber() int { return 7 }
func readFile(path string) (string, error) { return "", nil }
func note(a string, b int) { }`)
	if len(fs) != 3 {
		t.Fatalf("want 3 funcs, got %d", len(fs))
	}
	if fs[0].Name != "luckyNumber" || fs[0].Params != 0 || fs[0].Ret != "int" || fs[0].Fallible {
		t.Errorf("luckyNumber parsed wrong: %+v", fs[0])
	}
	if fs[1].Name != "readFile" || fs[1].Params != 1 || fs[1].Ret != "string" || !fs[1].Fallible {
		t.Errorf("readFile parsed wrong: %+v", fs[1])
	}
	if fs[2].Params != 2 || fs[2].Ret != "" {
		t.Errorf("note parsed wrong: %+v", fs[2])
	}
}
