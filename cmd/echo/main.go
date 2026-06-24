package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	noNewline = flag.Bool("n", false, "do not output the trailing newline")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Display a line of text.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	out := strings.Join(args, " ")
	if *noNewline {
		fmt.Print(out)
	} else {
		fmt.Println(out)
	}
	return nil
}
