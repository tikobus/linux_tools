package main

import (
	"compress/gzip"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		decompress = flag.Bool("d", false, "decompress")
		force      = flag.Bool("f", false, "force overwrite")
	)
	app := cliutil.NewApp("gzip", "Compress or decompress files using gzip.")
	flag.Usage = app.Usage
	flag.Parse()

	name := cliutil.ProgramName()
	if name == "gunzip" {
		*decompress = true
	}

	if flag.NArg() == 0 {
		if *decompress {
			return gunzip("-", "-", *force)
		}
		return gzipFile("-", "-", *force)
	}
	code := common.ExitSuccess
	for _, arg := range flag.Args() {
		if *decompress {
			if c := gunzip(arg, "", *force); c != 0 {
				code = c
			}
		} else {
			if c := gzipFile(arg, "", *force); c != 0 {
				code = c
			}
		}
	}
	return code
}

func gzipFile(src, dst string, force bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("gzip: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else {
			dst = src + ".gz"
		}
	}
	if !force && dst != "-" {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("gzip: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if err != nil {
		common.Warn("gzip: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	w := gzip.NewWriter(out)
	if _, err := io.Copy(w, in); err != nil {
		common.Warn("gzip: %v", err)
		return common.ExitFailure
	}
	w.Close()
	if src != "-" && dst != "-" {
		os.Remove(src)
	}
	return common.ExitSuccess
}

func gunzip(src, dst string, force bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("gunzip: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else if strings.HasSuffix(src, ".gz") {
			dst = strings.TrimSuffix(src, ".gz")
		} else {
			common.Warn("gunzip: %s: unknown suffix", src)
			return common.ExitFailure
		}
	}
	if !force && dst != "-" {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("gunzip: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if err != nil {
		common.Warn("gunzip: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	r, err := gzip.NewReader(in)
	if err != nil {
		common.Warn("gunzip: %v", err)
		return common.ExitFailure
	}
	defer r.Close()
	if _, err := io.Copy(out, r); err != nil {
		common.Warn("gunzip: %v", err)
		return common.ExitFailure
	}
	if src != "-" && dst != "-" {
		os.Remove(src)
	}
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
