package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	countOpt = flag.Bool("c", false, "prefix lines by the number of occurrences")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Report or omit repeated lines.")
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

	for _, file := range files {
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
	reader := bufio.NewReader(f)
	var prev string
	var count int
	first := true

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if line == "" && err == io.EOF {
			break
		}
		// Remove trailing newline for comparison
		trimmed := line
		if len(trimmed) > 0 && trimmed[len(trimmed)-1] == '\n' {
			trimmed = trimmed[:len(trimmed)-1]
		}

		if first {
			prev = trimmed
			count = 1
			first = false
		} else if trimmed == prev {
			count++
		} else {
			printLine(prev, count)
			prev = trimmed
			count = 1
		}
		if err == io.EOF {
			break
		}
	}
	if !first {
		printLine(prev, count)
	}
	return nil
}

func printLine(line string, count int) {
	if *countOpt {
		fmt.Printf("%7d %s\n", count, line)
	} else {
		fmt.Println(line)
	}
}
