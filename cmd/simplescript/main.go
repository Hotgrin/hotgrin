// Command simplescript is the friendly front door to the whole toolchain.
//
//	simplescript run     hello.ss      run a program
//	simplescript build   hello.ss      build a standalone program
//	simplescript build --windows x.ss  build a Windows .exe
//	simplescript check   hello.ss      check a program for problems
//	simplescript reveal  hello.ss      show the Go a program turns into
//	simplescript help                  show help
//
// It runs the Watcher before running or building, so a beginner sees friendly
// SimpleScript messages — never raw Go errors.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hotgrin/simplescript/internal/ast"
	"github.com/hotgrin/simplescript/internal/loader"
	"github.com/hotgrin/simplescript/internal/transpiler"
	"github.com/hotgrin/simplescript/internal/watcher"
)

const version = "SimpleScript 0.1 (clean rebuild)"

func main() {
	af := false
	windows := false
	var rest []string
	for _, a := range os.Args[1:] {
		switch a {
		case "--af", "-af", "--afrikaans":
			af = true
		case "--windows", "-windows", "--exe":
			windows = true
		default:
			rest = append(rest, a)
		}
	}

	lang := "en"
	if af {
		lang = "af"
	}

	if len(rest) == 0 {
		printHelp()
		return
	}

	cmd := rest[0]
	file := ""
	if len(rest) > 1 {
		file = rest[1]
	}
	var progArgs []string
	if len(rest) > 2 {
		progArgs = rest[2:]
	}

	switch cmd {
	case "help", "--help", "-h":
		printHelp()
	case "version", "--version", "-v":
		fmt.Println(version)
	case "check":
		cmdCheck(file, lang)
	case "run":
		cmdRun(file, lang, progArgs)
	case "test":
		cmdTest(file, lang)
	case "build":
		cmdBuild(file, lang, windows)
	case "reveal":
		cmdReveal(file, lang)
	default:
		fmt.Fprintf(os.Stderr, "I don't know the command %q. Try: simplescript help\n", cmd)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Print(`SimpleScript - a programming language that reads like plain English.

Usage:
  simplescript run     <file.ss>     Run a program (extra --flags pass to it)
  simplescript test    <file.ss>     Run the tests in a program
  simplescript build   <file.ss>     Build a standalone program you can share
  simplescript check   <file.ss>     Check a program for problems
  simplescript reveal  <file.ss>     Show the Go code a program becomes
  simplescript help                  Show this help
  simplescript version               Show the version

Options:
  --windows    With 'build', make a Windows .exe
  --af         Show messages in Afrikaans

Examples:
  simplescript run hello.ss
  simplescript build --windows hello.ss
  simplescript check --af hello.ss
`)
}

// load reads, lexes, and parses a file. It exits with a friendly message on a
// missing file or parse problems.
func load(file string) *ast.Program {
	if file == "" {
		fmt.Fprintln(os.Stderr, "Please tell me which file, e.g.  simplescript run hello.ss")
		os.Exit(1)
	}
	if _, err := os.Stat(file); err != nil {
		fmt.Fprintf(os.Stderr, "I couldn't open %q. Is the name right?\n", file)
		os.Exit(1)
	}
	prog, errs := loader.LoadFile(file)
	if len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "I couldn't understand part of your program:")
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  "+e)
		}
		os.Exit(1)
	}
	return prog
}

// report runs the Watcher and prints findings. It returns true if there were
// any errors (problems that must be fixed before running or building).
func report(prog *ast.Program, lang string) bool {
	findings := watcher.New(prog).Check()
	hasError := false
	for _, f := range findings {
		label := "idea   "
		switch f.Severity {
		case watcher.Error:
			label = "error  "
			hasError = true
		case watcher.Warning:
			label = "warning"
		}
		fmt.Fprintf(os.Stderr, "  %s line %d: %s\n", label, f.Line, f.Message(lang))
	}
	return hasError
}

func cmdCheck(file, lang string) {
	prog := load(file)
	if report(prog, lang) {
		os.Exit(1)
	}
	if len(watcher.New(prog).Check()) == 0 {
		fmt.Println("All good - I found no problems.")
	}
}

func cmdReveal(file, lang string) {
	prog := load(file)
	goSrc, _, _ := transpiler.New(prog).Transpile()
	fmt.Print(goSrc)
}

func cmdRun(file, lang string, progArgs []string) {
	prog := load(file)
	if report(prog, lang) {
		fmt.Fprintln(os.Stderr, "\nI found problems above, so I didn't run it. Fix those and try again.")
		os.Exit(1)
	}
	goSrc, _, _ := transpiler.New(prog).Transpile()
	dir := tempModule(goSrc)
	defer os.RemoveAll(dir)

	if !haveGo() {
		os.Exit(1)
	}
	runArgs := append([]string{"run", "."}, progArgs...)
	c := exec.Command("go", runArgs...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		os.Exit(1)
	}
}

func cmdTest(file, lang string) {
	prog := load(file)
	if report(prog, lang) {
		fmt.Fprintln(os.Stderr, "\nI found problems above, so I didn't run the tests.")
		os.Exit(1)
	}
	mainSrc, testSrc, _ := transpiler.New(prog).Transpile()
	if testSrc == "" {
		fmt.Println("No tests found. Add a 'test \"...\" ... end test' block.")
		return
	}
	if !haveGo() {
		os.Exit(1)
	}
	dir, _ := os.MkdirTemp("", "sstest")
	defer os.RemoveAll(dir)
	_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainSrc), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "main_test.go"), []byte(testSrc), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module ssprogram\n\ngo 1.22\n"), 0o644)

	c := exec.Command("go", "test", "-v", ".")
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		os.Exit(1)
	}
}

func cmdBuild(file, lang string, windows bool) {
	prog := load(file)
	if report(prog, lang) {
		fmt.Fprintln(os.Stderr, "\nI found problems above, so I didn't build. Fix those and try again.")
		os.Exit(1)
	}
	goSrc, _, _ := transpiler.New(prog).Transpile()
	dir := tempModule(goSrc)
	defer os.RemoveAll(dir)

	if !haveGo() {
		os.Exit(1)
	}

	base := strings.TrimSuffix(filepath.Base(file), ".ss")
	out := base
	if windows {
		out += ".exe"
	}
	outAbs, _ := filepath.Abs(out)

	c := exec.Command("go", "build", "-o", outAbs, ".")
	c.Dir = dir
	if windows {
		c.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	}
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "The build failed unexpectedly.")
		os.Exit(1)
	}
	fmt.Printf("Done. Built: %s\n", out)
	if !windows {
		fmt.Printf("Run it with:  ./%s\n", out)
	}
}

func tempModule(goSrc string) string {
	dir, err := os.MkdirTemp("", "simplescript")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte(goSrc), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module ssprogram\n\ngo 1.22\n"), 0o644)
	return dir
}

func haveGo() bool {
	if _, err := exec.LookPath("go"); err != nil {
		fmt.Fprintln(os.Stderr, "To run or build programs, SimpleScript needs Go installed for now.")
		fmt.Fprintln(os.Stderr, "Get it from https://go.dev/dl/ , then try again.")
		return false
	}
	return true
}
