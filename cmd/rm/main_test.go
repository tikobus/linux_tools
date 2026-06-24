package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestRmFile(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "a.txt", "hello")

	*recursiveFlag = false
	*forceFlag = false
	*interactiveFlag = false
	if err := run([]string{f}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(f); !os.IsNotExist(err) {
		t.Fatal("file should be removed")
	}
}

func TestRmRecursive(t *testing.T) {
	dir := testutil.TempDir(t)
	sub := filepath.Join(dir, "sub")
	os.Mkdir(sub, 0755)
	testutil.WriteFile(t, sub, "nested.txt", "data")

	*recursiveFlag = true
	*forceFlag = false
	*interactiveFlag = false
	if err := run([]string{sub}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(sub); !os.IsNotExist(err) {
		t.Fatal("directory should be removed")
	}
}

func TestRmForceNonExistent(t *testing.T) {
	dir := testutil.TempDir(t)
	f := filepath.Join(dir, "nonexistent")

	*recursiveFlag = false
	*forceFlag = true
	*interactiveFlag = false
	if err := run([]string{f}); err != nil {
		t.Fatal(err)
	}
}
