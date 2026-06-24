package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	numberFlag       = flag.Bool("n", false, "number all output lines")
	numberNonblankFlag = flag.Bool("b", false, "number nonempty output lines")
)

func main() {
	app := cliutil.NewApp("cat", "Concatenate FILE(s) to standard output.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"-"}
	}

	if err := run(args, os.Stdout); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string, out io.Writer) error {
	lineNum := 1
	for _, arg := range args {
		f, err := common.OpenInput(arg)
		if err != nil {
			return err
		}
		if arg != "-" {
			defer f.Close()
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if *numberNonblankFlag && strings.TrimSpace(line) != "" {
				fmt.Fprintf(out, "%6d\t%s\n", lineNum, line)
				lineNum++
			} else if *numberFlag {
				fmt.Fprintf(out, "%6d\t%s\n", lineNum, line)
				lineNum++
			} else {
				fmt.Fprintln(out, line)
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
