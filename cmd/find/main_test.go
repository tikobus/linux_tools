package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestFindBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, "a.txt", "hello")
	testutil.WriteFile(t, dir, "b.go", "package")

	var buf bytes.Buffer
	*nameFlag = "*.txt"
	*typFlag = ""
	*sizeFlag = ""
	if err := run(dir, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "a.txt") {
		t.Fatalf("expected a.txt in output, got: %s", out)
	}
	if strings.Contains(out, "b.go") {
		t.Fatalf("did not expect b.go in output, got: %s", out)
	}
}

func TestFindType(t *testing.T) {
	dir := testutil.TempDir(t)
	os.Mkdir(filepath.Join(dir, "sub"), 0755)
	testutil.WriteFile(t, dir, "file.txt", "data")

	var buf bytes.Buffer
	*nameFlag = ""
	*typFlag = "d"
	*sizeFlag = ""
	if err := run(dir, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "sub") {
		t.Fatalf("expected sub in output, got: %s", out)
	}
	if strings.Contains(out, "file.txt") {
		t.Fatalf("did not expect file.txt in output, got: %s", out)
	}
}

func TestFindSize(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, "small.txt", "x")
	testutil.WriteFile(t, dir, "big.txt", strings.Repeat("x", 2000))

	var buf bytes.Buffer
	*nameFlag = ""
	*typFlag = ""
	*sizeFlag = "+1k"
	if err := run(dir, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "big.txt") {
		t.Fatalf("expected big.txt in output, got: %s", out)
	}
	if strings.Contains(out, "small.txt") {
		t.Fatalf("did not expect small.txt in output, got: %s", out)
	}
}
