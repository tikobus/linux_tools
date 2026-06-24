package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestBzip2RoundTrip(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "data.txt", "hello bzip2 world")
	bz := filepath.Join(dir, "data.txt.bz2")

	os.Args = []string{"bzip2", src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("bzip2 compress failed: %d", code)
	}
	if _, err := os.Stat(bz); err != nil {
		t.Fatalf("bzip2 output missing: %v", err)
	}

	os.Args = []string{"bzip2", "-d", bz}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("bzip2 decompress failed: %d", code)
	}
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte("hello bzip2 world")) {
		t.Fatalf("round-trip mismatch: %s", data)
	}
}
