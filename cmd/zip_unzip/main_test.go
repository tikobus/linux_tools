package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestZipRoundTrip(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "data.txt", "hello zip world")
	archive := filepath.Join(dir, "test.zip")

	os.Args = []string{"zip", archive, src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("zip create failed: %d", code)
	}

	os.Args = []string{"unzip", archive}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("unzip failed: %d", code)
	}

	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte("hello zip world")) {
		t.Fatalf("round-trip mismatch: %s", data)
	}
}

func TestZipList(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "a.txt", "a")
	archive := filepath.Join(dir, "list.zip")

	os.Args = []string{"zip", archive, src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("zip create failed: %d", code)
	}

	os.Args = []string{"zip", "-l", archive}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("zip list failed: %d", code)
	}
}
