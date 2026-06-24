package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	recursiveFlag = flag.Bool("r", false, "copy directories recursively")
	verboseFlag = flag.Bool("v", false, "explain what is being done")
	interactiveFlag = flag.Bool("i", false, "prompt before overwrite")
)

func main() {
	app := cliutil.NewApp("cp", "Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.")
	flag.Usage = app.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		app.Usage()
		os.Exit(1)
	}

	if err := run(args); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	srcs := args[:len(args)-1]
	dest := args[len(args)-1]

	destInfo, err := os.Stat(dest)
	if len(srcs) > 1 || (err == nil && destInfo.IsDir()) {
		for _, src := range srcs {
			if err := copyPath(src, filepath.Join(dest, filepath.Base(src))); err != nil {
				return err
			}
		}
		return nil
	}

	if len(srcs) != 1 {
		return fmt.Errorf("missing destination file operand after %q", srcs[0])
	}
	return copyPath(srcs[0], dest)
}

func copyPath(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		if !*recursiveFlag {
			return fmt.Errorf("-r not specified; omitting directory %q", src)
		}
		return copyDir(src, dest)
	}

	return copyFile(src, dest, srcInfo)
}

func copyFile(src, dest string, srcInfo os.FileInfo) error {
	if *interactiveFlag {
		if _, err := os.Stat(dest); err == nil {
			fmt.Fprintf(os.Stderr, "overwrite %s? ", dest)
			reader := bufio.NewReader(os.Stdin)
			resp, _ := reader.ReadString('\n')
			if len(resp) == 0 || (resp[0] != 'y' && resp[0] != 'Y') {
				return nil
			}
		}
	}

	if *verboseFlag {
		fmt.Printf("%s -> %s\n", src, dest)
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func copyDir(src, dest string) error {
	if *verboseFlag {
		fmt.Printf("%s -> %s\n", src, dest)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dest, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())
		if err := copyPath(srcPath, destPath); err != nil {
			return err
		}
	}
	return nil
}
