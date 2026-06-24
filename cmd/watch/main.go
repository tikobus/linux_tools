package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		interval = flag.Duration("n", 2*time.Second, "seconds between updates")
	)
	app := cliutil.NewApp("watch", "Execute a program periodically.")
	flag.Usage = app.Usage
	flag.Parse()

	if flag.NArg() < 1 {
		common.Warn("watch: missing command")
		app.Usage()
		return common.ExitFailure
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return common.ExitSuccess
		default:
		}
		// clear screen
		fmt.Print("\033[H\033[2J")
		cmd := exec.CommandContext(ctx, flag.Arg(0), flag.Args()[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		time.Sleep(*interval)
	}
}

func main() {
	os.Exit(run())
}
