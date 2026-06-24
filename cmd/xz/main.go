package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/ulikunitz/xz"
	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		decompress = flag.Bool("d", false, "decompress")
		stdout     = flag.Bool("c", false, "write to stdout")
	)
	app := cliutil.NewApp("xz", "Compress or decompress files using xz.")
	flag.Usage = app.Usage
	flag.Parse()

	if flag.NArg() == 0 {
		if *decompress {
			return xzDecompress("-", "-", *stdout)
		}
		return xzCompress("-", "-", *stdout)
	}
	code := common.ExitSuccess
	for _, arg := range flag.Args() {
		if *decompress {
			if c := xzDecompress(arg, "", *stdout); c != 0 {
				code = c
			}
		} else {
			if c := xzCompress(arg, "", *stdout); c != 0 {
				code = c
			}
		}
	}
	return code
}

func xzCompress(src, dst string, stdout bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else {
			dst = src + ".xz"
		}
	}
	if dst != "-" && !stdout {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("xz: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if stdout {
		out = os.Stdout
	}
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	w, err := xz.NewWriter(out)
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	if _, err := io.Copy(w, in); err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	if err := w.Close(); err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	if src != "-" && dst != "-" && !stdout {
		os.Remove(src)
	}
	return common.ExitSuccess
}

func xzDecompress(src, dst string, stdout bool) int {
	in, err := common.OpenInput(src)
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	defer in.Close()
	if dst == "" {
		if src == "-" {
			dst = "-"
		} else if strings.HasSuffix(src, ".xz") {
			dst = strings.TrimSuffix(src, ".xz")
		} else {
			common.Warn("xz: %s: unknown suffix", src)
			return common.ExitFailure
		}
	}
	if dst != "-" && !stdout {
		if _, err := os.Stat(dst); err == nil {
			common.Warn("xz: %s already exists", dst)
			return common.ExitFailure
		}
	}
	out, err := common.CreateOutput(dst)
	if stdout {
		out = os.Stdout
	}
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	defer out.Close()
	r, err := xz.NewReader(in)
	if err != nil {
		common.Warn("xz: %v", err)
		return common.ExitFailure
	}
	if _, err := io.Copy(out, r); err != nil {
		common.Warn("xz: %v", err)
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
