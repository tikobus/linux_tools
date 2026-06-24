package cliutil

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// App holds command metadata.
type App struct {
	Name string
	Desc string
}

// NewApp creates a new App.
func NewApp(name, desc string) *App {
	return &App{Name: name, Desc: desc}
}

// Usage prints the standard help message.
func (a *App) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARG]...\n", a.Name)
	fmt.Fprintf(os.Stderr, "\n%s\n", a.Desc)
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
}

// ShowHelp prints help and exits 0.
func (a *App) ShowHelp() {
	a.Usage()
	os.Exit(0)
}

// ShowHelp prints help for the current command and exits 0.
func ShowHelp(name, desc string) {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [ARG]...\n", name)
	fmt.Fprintf(os.Stderr, "\n%s\n", desc)
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
	os.Exit(0)
}

// ProgramName returns the base name of os.Args[0].
func ProgramName() string {
	return filepath.Base(os.Args[0])
}

// Exit exits with the given code.
func Exit(code int) {
	os.Exit(code)
}
