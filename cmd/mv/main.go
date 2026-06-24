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
	interactiveFlag = flag.Bool("i", false, "prompt before overwrite")
	verboseFlag = flag.Bool("v", false, "explain what is being done")
)

func main() {
	app := cliutil.NewApp("mv", "Move (rename) SOURCE to DEST.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		app.Usage()
		os.Exit(1)
	}

	if err := run(args); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	srcs := args[:len(args)-1]
	dest := args[len(args)-1]

	destInfo, err := os.Stat(dest)
	if len(srcs) > 1 || (err == nil && destInfo.IsDir()) {
		for _, src := range srcs {
			if err := move(src, filepath.Join(dest, filepath.Base(src))); err != nil {
				return err
			}
		}
		return nil
	}

	if len(srcs) != 1 {
		return fmt.Errorf("missing destination file operand after %q", srcs[0])
	}
	return move(srcs[0], dest)
}

func move(src, dest string) error {
	if *interactiveFlag {
		if _, err := os.Stat(dest); err == nil {
			fmt.Fprintf(os.Stderr, "overwrite %s? ", dest)
			reader := bufio.NewReader(os.Stdin)
			resp, _ := reader.ReadString('\n')
			if len(resp) == 0 || (resp[0] != 'y' && resp[0] != 'Y') {
				return nil
			}
		}
	}

	if *verboseFlag {
		fmt.Printf("%s -> %s\n", src, dest)
	}

	return os.Rename(src, dest)
}
