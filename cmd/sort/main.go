package main

import (
	"flag"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	uniqueOpt = flag.Bool("u", false, "output only unique lines")
	numeric   = flag.Bool("n", false, "compare according to string numerical value")
	dictSort  = flag.Bool("d", false, "consider only blanks and alphanumeric characters")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Sort lines of text files.")
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

	var allLines []string
	for _, file := range files {
		f, err := common.OpenInput(file)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			return err
		}
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			if i == len(lines)-1 && line == "" {
				continue
			}
			allLines = append(allLines, line)
		}
	}

	sort.Slice(allLines, func(i, j int) bool {
		return less(allLines[i], allLines[j])
	})

	if *uniqueOpt {
		allLines = uniqueLines(allLines)
	}

	for _, line := range allLines {
		fmt.Println(line)
	}
	return nil
}

func less(a, b string) bool {
	if *numeric {
		na, erra := strconv.ParseFloat(strings.TrimSpace(a), 64)
		nb, errb := strconv.ParseFloat(strings.TrimSpace(b), 64)
		if erra == nil && errb == nil {
			return na < nb
		}
	}
	if *dictSort {
		a = dictKey(a)
		b = dictKey(b)
	}
	return a < b
}

func dictKey(s string) string {
	var out []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == ' ' || r == '\t' {
			out = append(out, r)
		}
	}
	return string(out)
}

func uniqueLines(lines []string) []string {
	var out []string
	for i, line := range lines {
		if i == 0 || line != lines[i-1] {
			out = append(out, line)
		}
	}
	return out
}
