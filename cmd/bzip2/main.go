package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/dsnet/compress/bzip2"
	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		decompress = flag.Bool("d", false, "decompress")
		stdout     = flag.Bool("c", false, "write to stdout")
	)
	app := cliutil.NewApp("bzip2", "Compress or decompress files using bzip2.")
	flag.Usage = app.Usage
	flag.Parse()

	if flag.NArg() == 0 {
		if *decompress {
			return bz2Decompress("-", "-", *stdout)
		}
		return bz2Compress("-", "-", *stdout)
	}
	code := common.ExitSuccess
	for _, arg := range flag.Args() {
		if *decompress {
			if c := bz2Decompress(arg, "", *stdout); c != 0 {
				code = c
			}
		} else {
			if c := bz2Compress(arg, "", *stdout); c != 0 {
				code = c
			}
		}
	}
	return code
}

func bz2Compress(src, dst string, stdout bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else {
			dst = src + ".bz2"
		}
	}
	if dst != "-" && !stdout {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("bzip2: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if stdout {
		out = os.Stdout
	}
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	w, err := bzip2.NewWriter(out, &bzip2.WriterConfig{Level: 6})
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	if _, err := io.Copy(w, in); err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	if err := w.Close(); err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	if src != "-" && dst != "-" && !stdout {
		os.Remove(src)
	}
	return common.ExitSuccess
}

func bz2Decompress(src, dst string, stdout bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else if strings.HasSuffix(src, ".bz2") {
			dst = strings.TrimSuffix(src, ".bz2")
		} else {
			common.Warn("bzip2: %s: unknown suffix", src)
			return common.ExitFailure
		}
	}
	if dst != "-" && !stdout {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("bzip2: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if stdout {
		out = os.Stdout
	}
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	r, err := bzip2.NewReader(in, nil)
	if err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	if _, err := io.Copy(out, r); err != nil {
		common.Warn("bzip2: %v", err)
		return common.ExitFailure
	}
	if src != "-" && dst != "-" && !stdout {
		os.Remove(src)
	}
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
