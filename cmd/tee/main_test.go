package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestTeeBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	outfile := dir + "/out.txt"

	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("hello world\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	stdout := captureStdout(t, func() {
		if err := run([]string{outfile}); err != nil {
			t.Fatal(err)
		}
	})
	if stdout != "hello world\n" {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
	data := testutil.ReadFile(t, outfile)
	if data != "hello world\n" {
		t.Fatalf("unexpected file content: %q", data)
	}
}

func TestTeeAppend(t *testing.T) {
	dir := testutil.TempDir(t)
	outfile := dir + "/out.txt"
	os.WriteFile(outfile, []byte("existing\n"), 0644)

	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("new line\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	*appendOpt = true
	defer func() { *appendOpt = false }()

	captureStdout(t, func() {
		if err := run([]string{outfile}); err != nil {
			t.Fatal(err)
		}
	})
	data := testutil.ReadFile(t, outfile)
	if data != "existing\nnew line\n" {
		t.Fatalf("unexpected file content: %q", data)
	}
}

func TestTeeMultiple(t *testing.T) {
	dir := testutil.TempDir(t)
	outfile1 := dir + "/out1.txt"
	outfile2 := dir + "/out2.txt"

	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("data\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	captureStdout(t, func() {
		if err := run([]string{outfile1, outfile2}); err != nil {
			t.Fatal(err)
		}
	})
	if testutil.ReadFile(t, outfile1) != "data\n" {
		t.Fatal("outfile1 mismatch")
	}
	if testutil.ReadFile(t, outfile2) != "data\n" {
		t.Fatal("outfile2 mismatch")
	}
}
