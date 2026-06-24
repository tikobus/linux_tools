package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestTarCreateExtract(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "hello.txt", "hello world")
	archive := filepath.Join(dir, "test.tar")

	os.Args = []string{"tar", "-c", "-f", archive, src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("tar create failed: %d", code)
	}

	extractDir := filepath.Join(dir, "extract")
	os.MkdirAll(extractDir, 0755)
	os.Args = []string{"tar", "-x", "-f", archive, "-C", extractDir}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("tar extract failed: %d", code)
	}
}

func TestTarList(t *testing.T) {
	dir := testutil.TempDir(t)
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	src := testutil.WriteFile(t, dir, "a.txt", "a")
	archive := filepath.Join(dir, "list.tar")

	os.Args = []string{"tar", "-c", "-f", archive, src}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("tar create failed: %d", code)
	}

	os.Args = []string{"tar", "-t", "-f", archive}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if code := run(); code != 0 {
		t.Fatalf("tar list failed: %d", code)
	}
}
