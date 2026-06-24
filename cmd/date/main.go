package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		format = flag.String("format", "", "output format (strftime style)")
		utc    = flag.Bool("u", false, "print UTC time")
	)
	app := cliutil.NewApp("date", "Print or set the system date and time.")
	flag.Usage = app.Usage
	flag.Parse()

	now := time.Now()
	if *utc {
		now = now.UTC()
	}

	f := *format
	if f == "" {
		f = time.RFC1123
	} else {
		f = strftime(f)
	}
	fmt.Println(now.Format(f))
	return common.ExitSuccess
}

func strftime(f string) string {
	replacements := map[string]string{
		"%Y": "2006",
		"%y": "06",
		"%m": "01",
		"%d": "02",
		"%H": "15",
		"%M": "04",
		"%S": "05",
		"%F": "2006-01-02",
		"%T": "15:04:05",
		"%D": "01/02/06",
		"%a": "Mon",
		"%A": "Monday",
		"%b": "Jan",
		"%B": "January",
		"%p": "PM",
		"%Z": "MST",
		"%z": "-0700",
	}
	for k, v := range replacements {
		f = strings.ReplaceAll(f, k, v)
	}
	return f
}

func main() {
	os.Exit(run())
}
