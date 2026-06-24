package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestXzRoundTrip(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "data.txt", "hello xz world")
	xzFile := filepath.Join(dir, "data.txt.xz")

	os.Args = []string{"xz", src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("xz compress failed: %d", code)
	}
	if _, err := os.Stat(xzFile); err != nil {
		t.Fatalf("xz output missing: %v", err)
	}

	os.Args = []string{"xz", "-d", xzFile}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("xz decompress failed: %d", code)
	}
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte("hello xz world")) {
		t.Fatalf("round-trip mismatch: %s", data)
	}
}
