package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	allFlag      = flag.Bool("a", false, "do not ignore entries starting with .")
	longFlag     = flag.Bool("l", false, "use a long listing format")
	humanFlag    = flag.Bool("h", false, "with -l, print human readable sizes")
	recursiveFlag = flag.Bool("R", false, "list subdirectories recursively")
)

func main() {
	app := cliutil.NewApp("ls", "List directory contents.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	if err := run(args, os.Stdout); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string, out io.Writer) error {
	multiple := len(args) > 1 || *recursiveFlag
	for i, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if *longFlag {
				fmt.Fprintln(out, formatLong(info, arg))
			} else {
				fmt.Fprintln(out, arg)
			}
			continue
		}
		if multiple {
			if i > 0 {
				fmt.Fprintln(out)
			}
			fmt.Fprintf(out, "%s:\n", arg)
		}
		if err := listDir(arg, out); err != nil {
			return err
		}
	}
	return nil
}

func listDir(dir string, out io.Writer) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var names []string
	for _, entry := range entries {
		name := entry.Name()
		if !*allFlag && strings.HasPrefix(name, ".") {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	if *longFlag {
		for _, name := range names {
			info, err := entryInfo(dir, name)
			if err != nil {
				continue
			}
			fmt.Fprintln(out, formatLong(info, name))
		}
	} else {
		for _, name := range names {
			fmt.Fprintln(out, name)
		}
	}

	if *recursiveFlag {
		for _, name := range names {
			path := filepath.Join(dir, name)
			info, err := os.Stat(path)
			if err != nil {
				continue
			}
			if info.IsDir() {
				fmt.Fprintf(out, "\n%s:\n", path)
				if err := listDir(path, out); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func entryInfo(dir, name string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(dir, name))
}

func formatLong(info os.FileInfo, name string) string {
	mode := info.Mode().String()
	size := strconv.FormatInt(info.Size(), 10)
	if *humanFlag {
		size = humanSize(info.Size())
	}
	mtime := info.ModTime().Format(time.RFC822)
	return fmt.Sprintf("%s %s %s %s", mode, size, mtime, name)
}

func humanSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
