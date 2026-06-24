package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		outPath = flag.String("O", "", "output file")
	)
	app := cliutil.NewApp("wget", "Download files from the web.")
	flag.Usage = app.Usage
	flag.Parse()

	if flag.NArg() < 1 {
		common.Warn("wget: missing URL")
		app.Usage()
		return common.ExitFailure
	}
	url := flag.Arg(0)
	resp, err := http.Get(url)
	if err != nil {
		common.Warn("wget: %v", err)
		return common.ExitFailure
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		common.Warn("wget: %s", resp.Status)
		return common.ExitFailure
	}

	out := *outPath
	if out == "" {
		out = filepath.Base(url)
		if out == "" || out == "." {
			out = "index.html"
		}
	}
	f, err := os.Create(out)
	if err != nil {
		common.Warn("wget: %v", err)
		return common.ExitFailure
	}
	defer f.Close()
	if _, err := io.Copy(f, resp.Body); err != nil {
		common.Warn("wget: %v", err)
		return common.ExitFailure
	}
	fmt.Println(out)
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
