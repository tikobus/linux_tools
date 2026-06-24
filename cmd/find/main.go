package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	nameFlag = flag.String("name", "", "match base name pattern")
	typFlag  = flag.String("type", "", "match file type (f=file, d=directory)")
	sizeFlag = flag.String("size", "", "match file size (e.g., +1k, -1M, 1G)")
)

func main() {
	app := cliutil.NewApp("find", "Search for files in a directory hierarchy.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	root := "."
	if len(args) > 0 {
		root = args[0]
	}

	if err := run(root, os.Stdout); err != nil {
		common.Fatal("%v", err)
	}
}

func run(root string, out io.Writer) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !match(info) {
			return nil
		}
		fmt.Fprintln(out, path)
		return nil
	})
}

func match(info os.FileInfo) bool {
	if *nameFlag != "" && !matchName(info.Name(), *nameFlag) {
		return false
	}
	if *typFlag != "" && !matchType(info, *typFlag) {
		return false
	}
	if *sizeFlag != "" && !matchSize(info, *sizeFlag) {
		return false
	}
	return true
}

func matchName(name, pattern string) bool {
	matched, _ := filepath.Match(pattern, name)
	return matched
}

func matchType(info os.FileInfo, typ string) bool {
	switch typ {
	case "f":
		return !info.IsDir()
	case "d":
		return info.IsDir()
	default:
		return false
	}
}

func matchSize(info os.FileInfo, sizeSpec string) bool {
	if len(sizeSpec) == 0 {
		return true
	}
	size := info.Size()
	spec := sizeSpec
	multiplier := int64(1)
	last := spec[len(spec)-1]
	switch last {
	case 'k', 'K':
		multiplier = 1024
		spec = spec[:len(spec)-1]
	case 'M':
		multiplier = 1024 * 1024
		spec = spec[:len(spec)-1]
	case 'G':
		multiplier = 1024 * 1024 * 1024
		spec = spec[:len(spec)-1]
	}

	val, err := strconv.ParseInt(spec, 10, 64)
	if err != nil {
		return false
	}
	val *= multiplier

	if strings.HasPrefix(sizeSpec, "+") {
		return size > val
	}
	if strings.HasPrefix(sizeSpec, "-") {
		return size < val
	}
	return size == val
}
