package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	deleteOpt = flag.String("d", "", "delete characters in SET1")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Translate or delete characters.")
	flag.Usage = app.Usage
	flag.Parse()

	if err := run(flag.Args()); err != nil {
		common.Fatal("%v", err)
	}
}

func run(args []string) error {
	if *deleteOpt != "" {
		return deleteChars(*deleteOpt)
	}
	if len(args) < 2 {
		return fmt.Errorf("tr: missing operand")
	}
	set1 := expand(args[0])
	set2 := expand(args[1])
	return translate(set1, set2)
}

func expand(s string) string {
	var out []rune
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+1] == '-' && s[i] < s[i+2] {
			for r := rune(s[i]); r <= rune(s[i+2]); r++ {
				out = append(out, r)
			}
			i += 2
		} else {
			out = append(out, rune(s[i]))
		}
	}
	return string(out)
}

func translate(set1, set2 string) error {
	m := make(map[rune]rune)
	runes1 := []rune(set1)
	runes2 := []rune(set2)
	for i, r := range runes1 {
		if i < len(runes2) {
			m[r] = runes2[i]
		} else {
			m[r] = runes2[len(runes2)-1]
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if nr, ok := m[r]; ok {
			os.Stdout.WriteString(string(nr))
		} else {
			os.Stdout.WriteString(string(r))
		}
	}
	return nil
}

func deleteChars(set string) error {
	set = expand(set)
	m := make(map[rune]bool)
	for _, r := range set {
		m[r] = true
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if !m[r] {
			os.Stdout.WriteString(string(r))
		}
	}
	return nil
}
