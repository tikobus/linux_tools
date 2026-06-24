package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestCpFile(t *testing.T) {
	dir := testutil.TempDir(t)
	src := testutil.WriteFile(t, dir, "src.txt", "hello")
	dest := filepath.Join(dir, "dest.txt")

	*recursiveFlag = false
	*verboseFlag = false
	*interactiveFlag = false
	if err := run([]string{src, dest}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Fatalf("expected 'hello', got %q", string(data))
	}
}

func TestCpRecursive(t *testing.T) {
	dir := testutil.TempDir(t)
	srcDir := filepath.Join(dir, "src")
	os.Mkdir(srcDir, 0755)
	testutil.WriteFile(t, srcDir, "file.txt", "data")
	destDir := filepath.Join(dir, "dest")

	*recursiveFlag = true
	*verboseFlag = false
	*interactiveFlag = false
	if err := run([]string{srcDir, destDir}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(destDir, "file.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "data" {
		t.Fatalf("expected 'data', got %q", string(data))
	}
}

func TestCpToDir(t *testing.T) {
	dir := testutil.TempDir(t)
	src := testutil.WriteFile(t, dir, "a.txt", "alpha")
	destDir := filepath.Join(dir, "d")
	os.Mkdir(destDir, 0755)

	*recursiveFlag = false
	*verboseFlag = false
	*interactiveFlag = false
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
