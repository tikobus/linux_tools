package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestMvFile(t *testing.T) {
	dir := testutil.TempDir(t)
	src := testutil.WriteFile(t, dir, "src.txt", "hello")
	dest := filepath.Join(dir, "dest.txt")

	*interactiveFlag = false
	*verboseFlag = false
	if err := run([]string{src, dest}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(src); !os.IsNotExist(err) {
		t.Fatal("source should not exist after move")
	}
	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Fatalf("expected 'hello', got %q", string(data))
	}
}

func TestMvToDir(t *testing.T) {
	dir := testutil.TempDir(t)
	src := testutil.WriteFile(t, dir, "a.txt", "alpha")
	destDir := filepath.Join(dir, "d")
	os.Mkdir(destDir, 0755)

	*interactiveFlag = false
	*verboseFlag = false
	if err := run([]string{src, destDir}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(destDir, "a.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "alpha" {
		t.Fatalf("expected 'alpha', got %q", string(data))
	}
}

func TestMvMultipleToDir(t *testing.T) {
	dir := testutil.TempDir(t)
	a := testutil.WriteFile(t, dir, "a.txt", "A")
	b := testutil.WriteFile(t, dir, "b.txt", "B")
	destDir := filepath.Join(dir, "d")
	os.Mkdir(destDir, 0755)

	*interactiveFlag = false
	*verboseFlag = false
	if err := run([]string{a, b, destDir}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(a); !os.IsNotExist(err) {
		t.Fatal("a.txt should not exist after move")
	}
	if _, err := os.Stat(b); !os.IsNotExist(err) {
		t.Fatal("b.txt should not exist after move")
	}
	data, _ := os.ReadFile(filepath.Join(destDir, "a.txt"))
	if string(data) != "A" {
		t.Fatalf("expected 'A', got %q", string(data))
	}
}
