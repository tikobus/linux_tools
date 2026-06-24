package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	linesOpt = flag.Int("n", 10, "print the first NUM lines")
	bytesOpt = flag.Int("c", -1, "print the first NUM bytes")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Output the first part of files.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	files := args
	if len(files) == 0 {
		files = []string{"-"}
	}

	multi := len(files) > 1

	for _, file := range files {
		if multi {
			fmt.Printf("==> %s <==\n", file)
		}
		f, err := common.OpenInput(file)
		if err != nil {
			return err
		}
		if err := process(f); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func process(f *os.File) error {
	if *bytesOpt >= 0 {
		_, err := io.CopyN(os.Stdout, f, int64(*bytesOpt))
		if err != nil && err != io.EOF {
			return err
		}
		return nil
	}

	// reader unused
	// Actually, let's just use a simple scanner approach
	// But scanner has a line limit. Let's use a simpler approach.
	buf := make([]byte, 1)
	lines := 0
	for {
		if lines >= *linesOpt {
			break
		}
		n, err := f.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf)
			if buf[0] == '\n' {
				lines++
			}
		}
		if err != nil {
			break
		}
	}
	return nil
}
