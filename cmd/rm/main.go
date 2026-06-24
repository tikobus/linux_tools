package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	recursiveFlag = flag.Bool("r", false, "remove directories and their contents recursively")
	forceFlag = flag.Bool("f", false, "ignore non-existent files, never prompt")
	interactiveFlag = flag.Bool("i", false, "prompt before every removal")
)

func main() {
	app := cliutil.NewApp("rm", "Remove files or directories.")
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
		if err := remove(arg); err != nil {
			if *forceFlag && os.IsNotExist(err) {
				continue
			}
			return err
		}
	}
	return nil
}

func remove(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		if !*recursiveFlag {
			return fmt.Errorf("%s: is a directory", path)
		}
		return removeDir(path)
	}

	if *interactiveFlag && !*forceFlag {
		fmt.Fprintf(os.Stderr, "remove %s? ", path)
		reader := bufio.NewReader(os.Stdin)
		resp, _ := reader.ReadString('\n')
		if len(resp) == 0 || (resp[0] != 'y' && resp[0] != 'Y') {
			return nil
		}
	}

	return os.Remove(path)
}

func removeDir(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		child := filepath.Join(path, entry.Name())
		if err := remove(child); err != nil {
			return err
		}
	}
	return os.Remove(path)
}
