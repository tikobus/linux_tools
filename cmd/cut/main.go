package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	bytesOpt  = flag.String("b", "", "select only these bytes")
	charsOpt  = flag.String("c", "", "select only these characters")
	fieldsOpt = flag.String("f", "", "select only these fields")
	delimOpt  = flag.String("d", "\t", "use DELIM instead of TAB for field delimiter")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Remove sections from each line of files.")
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

	var ranges [][2]int
	var mode string
	var delim string

	switch {
	case *bytesOpt != "":
		mode = "bytes"
		ranges = parseRanges(*bytesOpt)
	case *charsOpt != "":
		mode = "chars"
		ranges = parseRanges(*charsOpt)
	case *fieldsOpt != "":
		mode = "fields"
		ranges = parseRanges(*fieldsOpt)
		delim = *delimOpt
	default:
		return fmt.Errorf("you must specify -b, -c, or -f")
	}

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
			out := cutLine(line, mode, ranges, delim)
			if out != "" || mode == "fields" {
				fmt.Println(out)
			}
		}
	}
	return nil
}

func parseRanges(s string) [][2]int {
	var res [][2]int
	for _, part := range strings.Split(s, ",") {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			idx := strings.Index(part, "-")
			startStr := part[:idx]
			endStr := part[idx+1:]
			start := 1
			end := -1
			if startStr != "" {
				fmt.Sscanf(startStr, "%d", &start)
			}
			if endStr != "" {
				fmt.Sscanf(endStr, "%d", &end)
			}
			res = append(res, [2]int{start, end})
		} else {
			var n int
			fmt.Sscanf(part, "%d", &n)
			res = append(res, [2]int{n, n})
		}
	}
	return res
}

func cutLine(line, mode string, ranges [][2]int, delim string) string {
	switch mode {
	case "bytes":
		return cutBytes(line, ranges)
	case "chars":
		return cutChars(line, ranges)
	case "fields":
		return cutFields(line, ranges, delim)
	}
	return line
}

func cutBytes(line string, ranges [][2]int) string {
	bytes := []byte(line)
	var out []byte
	for _, r := range ranges {
		start := r[0] - 1
		end := r[1]
		if end == -1 {
			end = len(bytes)
		}
		if start < 0 {
			start = 0
		}
		if end > len(bytes) {
			end = len(bytes)
		}
		if start < end {
			out = append(out, bytes[start:end]...)
		}
	}
	return string(out)
}

func cutChars(line string, ranges [][2]int) string {
	runes := []rune(line)
	var out []rune
	for _, r := range ranges {
		start := r[0] - 1
		end := r[1]
		if end == -1 {
			end = len(runes)
		}
		if start < 0 {
			start = 0
		}
		if end > len(runes) {
			end = len(runes)
		}
		if start < end {
			out = append(out, runes[start:end]...)
		}
	}
	return string(out)
}

func cutFields(line string, ranges [][2]int, delim string) string {
	fields := strings.Split(line, delim)
	var out []string
	for _, r := range ranges {
		start := r[0] - 1
		end := r[1]
		if end == -1 {
			end = len(fields)
		}
		if start < 0 {
			start = 0
		}
		if end > len(fields) {
			end = len(fields)
		}
		for i := start; i < end && i < len(fields); i++ {
			out = append(out, fields[i])
		}
	}
	return strings.Join(out, delim)
}

func countRunes(s string) int {
	return utf8.RuneCountInString(s)
}
