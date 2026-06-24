package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		all    = flag.Bool("a", false, "print all information")
		sysname = flag.Bool("s", false, "print the kernel name")
		nodename = flag.Bool("n", false, "print the network node hostname")
		release = flag.Bool("r", false, "print the kernel release")
		version = flag.Bool("v", false, "print the kernel version")
		machine = flag.Bool("m", false, "print the machine hardware name")
	)
	app := cliutil.NewApp("uname", "Print system information.")
	flag.Usage = app.Usage
	flag.Parse()

	if *all {
		fmt.Printf("%s %s %s %s %s\n", runtime.GOOS, getHostname(), runtime.Version(), runtime.Version(), runtime.GOARCH)
		return common.ExitSuccess
	}

	if !*sysname && !*nodename && !*release && !*version && !*machine {
		*sysname = true
	}

	var out []string
	if *sysname {
		out = append(out, runtime.GOOS)
	}
	if *nodename {
		out = append(out, getHostname())
	}
	if *release {
		out = append(out, runtime.Version())
	}
	if *version {
		out = append(out, runtime.Version())
	}
	if *machine {
		out = append(out, runtime.GOARCH)
	}

	for i, s := range out {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(s)
	}
	fmt.Println()
	return common.ExitSuccess
}

func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}

func main() {
	os.Exit(run())
}
