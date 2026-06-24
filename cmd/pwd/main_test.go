package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPwd(t *testing.T) {
	var buf bytes.Buffer
	if err := run(&buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	wd, _ := os.Getwd()
	if !strings.Contains(out, wd) {
		t.Fatalf("expected %q in output, got: %s", wd, out)
	}
}

func TestPwdNotEmpty(t *testing.T) {
	var buf bytes.Buffer
	if err := run(&buf); err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Fatal("expected non-empty output")
	}
}
