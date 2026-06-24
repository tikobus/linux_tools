package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	linesOpt = flag.Int("n", 10, "output the last NUM lines")
	bytesOpt = flag.Int("c", -1, "output the last NUM bytes")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Output the last part of files.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	files := args
	if len(files) == 0 {
		files = []string{"-"}
	}

	multi := len(files) > 1

	for _, file := range files {
		if multi {
			fmt.Printf("==> %s <==\n", file)
		}
		f, err := common.OpenInput(file)
		if err != nil {
			return err
		}
		if err := process(f); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func process(f *os.File) error {
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	if *bytesOpt >= 0 {
		start := len(data) - *bytesOpt
		if start < 0 {
			start = 0
		}
		os.Stdout.Write(data[start:])
		return nil
	}

	lines := strings.Split(string(data), "\n")
	// Handle trailing newline
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	start := len(lines) - *linesOpt
	if start < 0 {
		start = 0
	}
	for i := start; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
	return nil
}
