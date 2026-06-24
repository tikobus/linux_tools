package main

import (
	"bytes"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestCatBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	path := testutil.WriteFile(t, dir, "file.txt", "hello\nworld\n")

	var buf bytes.Buffer
	*numberFlag = false
	*numberNonblankFlag = false
	if err := run([]string{path}, &buf); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "hello\nworld\n" {
		t.Fatalf("unexpected output: %q", buf.String())
	}
}

func TestCatNumber(t *testing.T) {
	dir := testutil.TempDir(t)
	path := testutil.WriteFile(t, dir, "file.txt", "a\nb\n")

	var buf bytes.Buffer
	*numberFlag = true
	*numberNonblankFlag = false
	if err := run([]string{path}, &buf); err != nil {
		t.Fatal(err)
	}
	expected := "     1\ta\n     2\tb\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}

func TestCatNumberNonblank(t *testing.T) {
	dir := testutil.TempDir(t)
	path := testutil.WriteFile(t, dir, "file.txt", "a\n\nb\n")

	var buf bytes.Buffer
	*numberFlag = false
	*numberNonblankFlag = true
	if err := run([]string{path}, &buf); err != nil {
		t.Fatal(err)
	}
	expected := "     1\ta\n\n     2\tb\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}
