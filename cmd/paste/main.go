package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	delimOpt = flag.String("d", "\t", "reuse characters from LIST instead of TABs")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Merge lines of files.")
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

	var readers [][]string
	maxLines := 0

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
		// Remove trailing empty line if file ends with newline
		if len(lines) > 0 && lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		readers = append(readers, lines)
		if len(lines) > maxLines {
			maxLines = len(lines)
		}
	}

	delim := *delimOpt
	if delim == "" {
		delim = "\t"
	}

	for i := 0; i < maxLines; i++ {
		var parts []string
		for _, lines := range readers {
			if i < len(lines) {
				parts = append(parts, lines[i])
			} else {
				parts = append(parts, "")
			}
		}
		fmt.Println(strings.Join(parts, delim))
	}
	return nil
}
