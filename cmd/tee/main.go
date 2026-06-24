package main

import (
	"flag"
	"io"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	appendOpt = flag.Bool("a", false, "append to the given FILEs, do not overwrite")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Read from standard input and write to standard output and files.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	files := args

	var writers []io.Writer
	writers = append(writers, os.Stdout)

	for _, file := range files {
		var f *os.File
		var err error
		if *appendOpt {
			f, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		} else {
			f, err = os.Create(file)
		}
		if err != nil {
			return err
		}
		defer f.Close()
		writers = append(writers, f)
	}

	mw := io.MultiWriter(writers...)
	_, err := io.Copy(mw, os.Stdin)
	return err
}
