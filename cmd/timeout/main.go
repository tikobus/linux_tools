package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	return runArgs(os.Args[1:])
}

func runArgs(args []string) int {
	fs := flag.NewFlagSet("timeout", flag.ContinueOnError)
	sig := fs.Int("s", 15, "signal to send on timeout")
	app := cliutil.NewApp("timeout", "Run a command with a time limit.")
	fs.Usage = app.Usage
	if err := fs.Parse(args); err != nil {
		return common.ExitFailure
	}
	if fs.NArg() < 2 {
		common.Warn("timeout: missing operand")
		app.Usage()
		return common.ExitFailure
	}
	d, err := time.ParseDuration(fs.Arg(0))
	if err != nil {
		common.Warn("timeout: invalid duration %q", fs.Arg(0))
		return common.ExitFailure
	}
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	cmd := exec.CommandContext(ctx, fs.Arg(1), fs.Args()[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			common.Warn("timeout: sending signal %d", *sig)
			if cmd.Process != nil {
				cmd.Process.Signal(syscall.Signal(*sig))
			}
			return 124
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		common.Warn("timeout: %v", err)
		return common.ExitFailure
	}
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
