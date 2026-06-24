package main

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestCalCurrentMonth(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cal"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestCalSpecificMonth(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cal", "3", "2024"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestCalYear(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cal", "2024"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestCalInvalidMonth(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cal", "13", "2024"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code == 0 {
		t.Fatal("expected non-zero exit for invalid month")
	}
}

func TestCalOutputContainsMonth(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cal", "1", "2024"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := run()
	w.Close()
	os.Stdout = oldStdout
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	output := string(buf[:n])
	if !strings.Contains(output, "January") {
		t.Fatalf("expected output to contain January, got: %s", output)
	}
}
