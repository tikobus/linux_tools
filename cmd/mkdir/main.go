package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	parentsFlag = flag.Bool("p", false, "make parent directories as needed")
	modeFlag = flag.String("m", "755", "set file mode (as octal number)")
)

func main() {
	app := cliutil.NewApp("mkdir", "Create directories.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		app.Usage()
		os.Exit(1)
	}

	mode, err := strconv.ParseUint(*modeFlag, 8, 32)
	if err != nil {
		common.Fatal("invalid mode: %v", err)
	}

	if err := run(args, os.FileMode(mode)); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string, mode os.FileMode) error {
	for _, arg := range args {
		if *parentsFlag {
			if err := os.MkdirAll(arg, mode); err != nil {
				return err
			}
		} else {
			if err := os.Mkdir(arg, mode); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
