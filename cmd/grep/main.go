package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	ignoreCase = flag.Bool("i", false, "ignore case distinctions")
	invert     = flag.Bool("v", false, "invert match")
	lineNum    = flag.Bool("n", false, "print line number")
	recursive  = flag.Bool("r", false, "recursive directory search")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Search for PATTERN in each FILE or standard input.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		app.Usage()
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]
	if len(files) == 0 {
		files = []string{"-"}
	}

	if err := run(pattern, files); err != nil {
		common.Fatal("%v", err)
	}
}

func run(pattern string, files []string) error {
	re, err := compilePattern(pattern)
	if err != nil {
		return err
	}

	multi := len(files) > 1 || *recursive

	for _, file := range files {
		if *recursive && isDir(file) {
			if err := grepDir(re, file, multi); err != nil {
				return err
			}
		} else {
			if err := grepFile(re, file, multi); err != nil {
				return err
			}
		}
	}
	return nil
}

func compilePattern(pattern string) (*regexp.Regexp, error) {
	if *ignoreCase {
		pattern = "(?i:" + pattern + ")"
	}
	return regexp.Compile(pattern)
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func grepDir(re *regexp.Regexp, dir string, multi bool) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return grepFile(re, path, true)
	})
}

func grepFile(re *regexp.Regexp, filename string, multi bool) error {
	f, err := common.OpenInput(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if line == "" && err == io.EOF {
			break
		}
		lineNo++
		// Remove trailing newline for matching, but preserve for output
		matchLine := strings.TrimSuffix(line, "\n")
		matched := re.MatchString(matchLine)
		if *invert {
			matched = !matched
		}
		if matched {
			prefix := ""
			if multi {
				prefix = filename + ":"
			}
			if *lineNum {
				prefix += fmt.Sprintf("%d:", lineNo)
			}
			fmt.Print(prefix + line)
		}
		if err == io.EOF {
			break
		}
	}
	return nil
}
