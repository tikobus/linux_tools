package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestMkdirBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "newdir")

	*parentsFlag = false
	*modeFlag = "755"
	if err := run([]string{d}, 0755); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(d)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatal("expected directory")
	}
}

func TestMkdirParents(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "a", "b", "c")

	*parentsFlag = true
	*modeFlag = "755"
	if err := run([]string{d}, 0755); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(d)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatal("expected directory")
	}
}

func TestMkdirMode(t *testing.T) {
	dir := testutil.TempDir(t)
	d := filepath.Join(dir, "modedir")

	*parentsFlag = false
	*modeFlag = "700"
	if err := run([]string{d}, 0700); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(d)
	if err != nil {
		t.Fatal(err)
	}
	perm := info.Mode().Perm()
	if perm != 0700 {
		t.Fatalf("expected mode 0700, got %04o", perm)
	}
}
