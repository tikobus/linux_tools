package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestLsBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, "a.txt", "hello")
	testutil.WriteFile(t, dir, "b.txt", "world")

	var buf bytes.Buffer
	*allFlag = false
	*longFlag = false
	*humanFlag = false
	*recursiveFlag = false
	if err := run([]string{dir}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "a.txt") || !strings.Contains(out, "b.txt") {
		t.Fatalf("expected a.txt and b.txt in output, got: %s", out)
	}
}

func TestLsAll(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, ".hidden", "secret")
	testutil.WriteFile(t, dir, "visible", "open")

	var buf bytes.Buffer
	*allFlag = true
	*longFlag = false
	*humanFlag = false
	*recursiveFlag = false
	if err := run([]string{dir}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, ".hidden") || !strings.Contains(out, "visible") {
		t.Fatalf("expected .hidden and visible in output, got: %s", out)
	}
}

func TestLsLong(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, "file.txt", "content")

	var buf bytes.Buffer
	*allFlag = false
	*longFlag = true
	*humanFlag = false
	*recursiveFlag = false
	if err := run([]string{dir}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "file.txt") {
		t.Fatalf("expected file.txt in output, got: %s", out)
	}
	if !strings.Contains(out, "-") {
		t.Fatalf("expected mode info in output, got: %s", out)
	}
}

func TestLsRecursive(t *testing.T) {
	dir := testutil.TempDir(t)
	sub := filepath.Join(dir, "sub")
	os.Mkdir(sub, 0755)
	testutil.WriteFile(t, sub, "nested.txt", "data")

	var buf bytes.Buffer
	*allFlag = false
	*longFlag = false
	*humanFlag = false
	*recursiveFlag = true
	if err := run([]string{dir}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "nested.txt") {
		t.Fatalf("expected nested.txt in output, got: %s", out)
	}
}
