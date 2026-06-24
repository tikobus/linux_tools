package main

import (
	"flag"
	"os"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func main() {
	app := cliutil.NewApp("touch", "Change file timestamps or create empty files.")
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
	now := time.Now()
	for _, arg := range args {
		if _, err := os.Stat(arg); os.IsNotExist(err) {
			f, err := os.Create(arg)
			if err != nil {
				return err
			}
			f.Close()
		} else if err != nil {
			return err
		} else {
			if err := os.Chtimes(arg, now, now); err != nil {
				return err
			}
		}
	}
	return nil
}
