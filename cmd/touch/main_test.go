package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestTouchCreate(t *testing.T) {
	dir := testutil.TempDir(t)
	f := filepath.Join(dir, "newfile")

	if err := run([]string{f}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(f); err != nil {
		t.Fatal(err)
	}
}

func TestTouchUpdate(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "existing", "data")
	before, _ := os.Stat(f)
	time.Sleep(10 * time.Millisecond)

	if err := run([]string{f}); err != nil {
		t.Fatal(err)
	}
	after, _ := os.Stat(f)
	if !after.ModTime().After(before.ModTime()) {
		t.Fatal("modtime should be updated")
	}
}

func TestTouchMultiple(t *testing.T) {
	dir := testutil.TempDir(t)
	a := filepath.Join(dir, "a")
	b := filepath.Join(dir, "b")

	if err := run([]string{a, b}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(a); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(b); err != nil {
		t.Fatal(err)
	}
}
