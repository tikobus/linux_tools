package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func main() {
	app := cliutil.NewApp("pwd", "Print name of current working directory.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(os.Stdout); err != nil {
		common.Fatal("%v", err)
	}
}

func run(out io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, wd)
	return nil
}
