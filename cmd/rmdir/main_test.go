package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestRmdirBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "empty")
	os.Mkdir(d, 0755)

	*parentsFlag = false
	if err := run([]string{d}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(d); !os.IsNotExist(err) {
		t.Fatal("directory should be removed")
	}
}

func TestRmdirParents(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "a", "b")
	os.MkdirAll(d, 0755)

	*parentsFlag = true
	if err := run([]string{d}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Fatal("ancestor should be removed if empty")
	}
}

func TestRmdirNonEmpty(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "nonempty")
	os.Mkdir(d, 0755)
	testutil.WriteFile(t, d, "file.txt", "data")

	*parentsFlag = false
	if err := run([]string{d}); err == nil {
		t.Fatal("expected error for non-empty directory")
	}
}
