package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func main() {
	app := cliutil.NewApp("realpath", "Print the resolved absolute path.")
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
	for _, arg := range args {
		abs, err := filepath.Abs(arg)
		if err != nil {
			return err
		}
		resolved, err := filepath.EvalSymlinks(abs)
		if err != nil {
			resolved = abs
		}
		fmt.Fprintln(out, resolved)
	}
	return nil
}
