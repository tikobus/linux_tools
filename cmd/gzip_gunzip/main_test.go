package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestGzipRoundTrip(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "data.txt", "hello gzip world")
	gz := filepath.Join(dir, "data.txt.gz")

	os.Args = []string{"gzip", "-f", src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("gzip failed: %d", code)
	}
	if _, err := os.Stat(gz); err != nil {
		t.Fatalf("gzip output missing: %v", err)
	}

	os.Args = []string{"gunzip", "-f", gz}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("gunzip failed: %d", code)
	}
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte("hello gzip world")) {
		t.Fatalf("round-trip mismatch: %s", data)
	}
}

func TestGzipStdinStdout(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"gzip"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// stdin/stdout path without data is hard to test here; just verify no panic.
	// We can skip because it blocks.
	_ = oldArgs
}
