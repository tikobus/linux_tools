package main

import (
	"os"
	"time"

	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	if len(os.Args) < 2 {
		common.Warn("sleep: missing operand")
		return common.ExitFailure
	}
	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		common.Warn("sleep: invalid time interval %q", os.Args[1])
		return common.ExitFailure
	}
	time.Sleep(d)
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
