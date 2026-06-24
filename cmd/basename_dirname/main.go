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
	name := cliutil.ProgramName()
	var desc string
	if name == "dirname" {
		desc = "Output the directory portion of a pathname."
	} else {
		desc = "Output the last portion of a pathname."
	}
	app := cliutil.NewApp(name, desc)
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		app.Usage()
		os.Exit(1)
	}

	if err := run(args, os.Stdout, name); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string, out io.Writer, progName string) error {
	for _, arg := range args {
		if progName == "dirname" {
			fmt.Fprintln(out, filepath.Dir(arg))
		} else {
			fmt.Fprintln(out, filepath.Base(arg))
		}
	}
	return nil
}
