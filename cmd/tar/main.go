package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		create = flag.Bool("c", false, "create archive")
		extract = flag.Bool("x", false, "extract archive")
		list = flag.Bool("t", false, "list contents")
		file = flag.String("f", "", "archive file")
		chdir = flag.String("C", "", "change to directory")
	)
	app := cliutil.NewApp("tar", "Create, extract, or list tar archives.")
	flag.Usage = app.Usage
	flag.Parse()

	if *chdir != "" {
		if err := os.Chdir(*chdir); err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
	}
	if *file == "" {
		common.Warn("tar: -f is required")
		return common.ExitFailure
	}

	if *create {
		return tarCreate(*file, flag.Args())
	}
	if *extract {
		return tarExtract(*file)
	}
	if *list {
		return tarList(*file)
	}
	common.Warn("tar: one of -c, -x, -t is required")
	return common.ExitFailure
}

func tarCreate(name string, files []string) int {
	f, err := os.Create(name)
	if err != nil {
		common.Warn("tar: %v", err)
		return common.ExitFailure
	}
	defer f.Close()
	w := tar.NewWriter(f)
	defer w.Close()
	for _, path := range files {
		info, err := os.Stat(path)
		if err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		hdr.Name = filepath.ToSlash(path)
		if err := w.WriteHeader(hdr); err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		if !info.IsDir() {
			fp, err := os.Open(path)
			if err != nil {
				common.Warn("tar: %v", err)
				return common.ExitFailure
			}
			if _, err := io.Copy(w, fp); err != nil {
				fp.Close()
				common.Warn("tar: %v", err)
				return common.ExitFailure
			}
			fp.Close()
		}
	}
	return common.ExitSuccess
}

func tarExtract(name string) int {
	f, err := os.Open(name)
	if err != nil {
		common.Warn("tar: %v", err)
		return common.ExitFailure
	}
	defer f.Close()
	r := tar.NewReader(f)
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		path := filepath.FromSlash(hdr.Name)
		if hdr.Typeflag == tar.TypeDir {
			os.MkdirAll(path, os.FileMode(hdr.Mode))
			continue
		}
		os.MkdirAll(filepath.Dir(path), 0755)
		out, err := os.Create(path)
		if err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		if _, err := io.Copy(out, r); err != nil {
			out.Close()
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		out.Close()
		os.Chmod(path, os.FileMode(hdr.Mode))
	}
	return common.ExitSuccess
}

func tarList(name string) int {
	f, err := os.Open(name)
	if err != nil {
		common.Warn("tar: %v", err)
		return common.ExitFailure
	}
	defer f.Close()
	r := tar.NewReader(f)
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			common.Warn("tar: %v", err)
			return common.ExitFailure
		}
		fmt.Println(hdr.Name)
	}
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
