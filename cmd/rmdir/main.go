package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	parentsFlag = flag.Bool("p", false, "remove DIRECTORY and its ancestors")
)

func main() {
	app := cliutil.NewApp("rmdir", "Remove empty directories.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		app.Usage()
		os.Exit(1)
	}

	if err := run(args); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	for _, arg := range args {
		if err := rmdir(arg); err != nil {
			return err
		}
	}
	return nil
}

func rmdir(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	if *parentsFlag {
		for {
			parent := filepath.Dir(path)
			if parent == path {
				break
			}
			if err := os.Remove(parent); err != nil {
				break
			}
			path = parent
		}
	}
	return nil
}
