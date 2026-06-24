package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		extract = flag.Bool("x", false, "extract")
		list    = flag.Bool("l", false, "list")
	)
	app := cliutil.NewApp("zip", "Package and compress files into a zip archive.")
	flag.Usage = app.Usage
	flag.Parse()

	name := cliutil.ProgramName()
	if name == "unzip" {
		*extract = true
	}

	if flag.NArg() < 1 {
		common.Warn("zip: missing archive")
		return common.ExitFailure
	}
	archive := flag.Arg(0)

	if *extract {
		return unzip(archive)
	}
	if *list {
		return listZip(archive)
	}
	return zipCreate(archive, flag.Args()[1:])
}

func zipCreate(archive string, files []string) int {
	f, err := os.Create(archive)
	if err != nil {
		common.Warn("zip: %v", err)
		return common.ExitFailure
	}
	defer f.Close()
	w := zip.NewWriter(f)
	defer w.Close()
	for _, path := range files {
		info, err := os.Stat(path)
		if err != nil {
			common.Warn("zip: %v", err)
			return common.ExitFailure
		}
		if info.IsDir() {
			filepath.Walk(path, func(p string, fi os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				name := filepath.ToSlash(p)
				if fi.IsDir() {
					name += "/"
				}
				hdr, err := zip.FileInfoHeader(fi)
				if err != nil {
					return err
				}
				hdr.Name = name
				if fi.IsDir() {
					_, err := w.CreateHeader(hdr)
					return err
				}
				fw, err := w.CreateHeader(hdr)
				if err != nil {
					return err
				}
				fp, err := os.Open(p)
				if err != nil {
					return err
				}
				defer fp.Close()
				_, err = io.Copy(fw, fp)
				return err
			})
		} else {
			fw, err := w.Create(filepath.ToSlash(path))
			if err != nil {
				common.Warn("zip: %v", err)
				return common.ExitFailure
			}
			fp, err := os.Open(path)
			if err != nil {
				common.Warn("zip: %v", err)
				return common.ExitFailure
			}
			defer fp.Close()
			if _, err := io.Copy(fw, fp); err != nil {
				common.Warn("zip: %v", err)
				return common.ExitFailure
			}
		}
	}
	return common.ExitSuccess
}

func unzip(archive string) int {
	r, err := zip.OpenReader(archive)
	if err != nil {
		common.Warn("unzip: %v", err)
		return common.ExitFailure
	}
	defer r.Close()
	for _, f := range r.File {
		path := filepath.FromSlash(f.Name)
		if strings.Contains(path, "..") {
			common.Warn("unzip: skipping malicious path: %s", f.Name)
			continue
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		}
		os.MkdirAll(filepath.Dir(path), 0755)
		out, err := os.Create(path)
		if err != nil {
			common.Warn("unzip: %v", err)
			return common.ExitFailure
		}
		rc, err := f.Open()
		if err != nil {
			out.Close()
			common.Warn("unzip: %v", err)
			return common.ExitFailure
		}
		if _, err := io.Copy(out, rc); err != nil {
			out.Close()
			rc.Close()
			common.Warn("unzip: %v", err)
			return common.ExitFailure
		}
		out.Close()
		rc.Close()
		os.Chmod(path, f.Mode())
	}
	return common.ExitSuccess
}

func listZip(archive string) int {
	r, err := zip.OpenReader(archive)
	if err != nil {
		common.Warn("zip: %v", err)
		return common.ExitFailure
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Println(f.Name)
	}
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
