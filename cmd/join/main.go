package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	field1Opt = flag.Int("1", 1, "join on this field of file 1")
	field2Opt = flag.Int("2", 1, "join on this field of file 2")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Join lines of two files on a common field.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("join: missing file arguments")
	}

	file1 := args[0]
	file2 := args[1]

	map1, err := readFile(file1, *field1Opt)
	if err != nil {
		return err
	}
	map2, err := readFile(file2, *field2Opt)
	if err != nil {
		return err
	}

	// Print matching lines
	for key, lines1 := range map1 {
		if lines2, ok := map2[key]; ok {
			for _, l1 := range lines1 {
				for _, l2 := range lines2 {
					fmt.Println(joinLines(l1, l2, key))
				}
			}
		}
	}
	return nil
}

func readFile(filename string, fieldNum int) (map[string][]string, error) {
	f, err := common.OpenInput(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make(map[string][]string)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if line == "" && err == io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		fields := strings.Fields(line)
		if fieldNum-1 < len(fields) {
			key := fields[fieldNum-1]
			result[key] = append(result[key], line)
		}
		if err == io.EOF {
			break
		}
	}
	return result, nil
}

func joinLines(line1, line2, key string) string {
	fields1 := strings.Fields(line1)
	fields2 := strings.Fields(line2)

	var out []string
	out = append(out, key)

	for _, f := range fields1 {
		if f != key {
			out = append(out, f)
		}
	}
	for _, f := range fields2 {
		if f != key {
			out = append(out, f)
		}
	}
	return strings.Join(out, " ")
}
