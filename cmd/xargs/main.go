package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	delimitOpt = flag.String("d", "\n", "items delimited by characters in string")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Build and execute command lines from standard input.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	cmd := "echo"
	var cmdArgs []string
	if len(args) > 0 {
		cmd = args[0]
		cmdArgs = args[1:]
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if line == "" && err == io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")

		allArgs := append([]string{}, cmdArgs...)
		allArgs = append(allArgs, line)
		c := exec.Command(cmd, allArgs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return err
		}
		if err == io.EOF {
			break
		}
	}
	return nil
}
