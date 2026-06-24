package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func main() {
	app := cliutil.NewApp("which", "Locate a command in the PATH.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		app.Usage()
		os.Exit(1)
	}

	if err := run(args, os.Stdout); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string, out io.Writer) error {
	pathEnv := os.Getenv("PATH")
	paths := filepath.SplitList(pathEnv)
	found := false
	for _, arg := range args {
		for _, dir := range paths {
			full := filepath.Join(dir, arg)
			if info, err := os.Stat(full); err == nil && !info.IsDir() {
				fmt.Fprintln(out, full)
				found = true
				break
			}
			// On Windows, also try with .exe extension
			if strings.Contains(full, ".") {
				continue
			}
			for _, ext := range []string{".exe", ".cmd", ".bat"} {
				fullExt := full + ext
				if info, err := os.Stat(fullExt); err == nil && !info.IsDir() {
					fmt.Fprintln(out, fullExt)
					found = true
					break
				}
			}
		}
	}
	if !found {
		return fmt.Errorf("no command found in PATH")
	}
	return nil
}
