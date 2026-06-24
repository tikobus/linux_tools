package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Format and print data.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("printf: missing format")
	}
	format := args[0]
	values := args[1:]

	// Handle escape sequences first
	out := strings.Replace(format, "\\n", "\n", -1)
	out = strings.Replace(out, "\\t", "\t", -1)

	// Replace format specifiers with values
	valIdx := 0
	for {
		idx := strings.Index(out, "%")
		if idx == -1 || idx == len(out)-1 {
			break
		}
		spec := out[idx+1]
		if spec == 's' || spec == 'd' || spec == 'f' {
			if valIdx < len(values) {
				out = out[:idx] + values[valIdx] + out[idx+2:]
				valIdx++
			} else {
				out = out[:idx] + out[idx+2:]
			}
		} else {
			break
		}
	}

	fmt.Print(out)
	return nil
}
