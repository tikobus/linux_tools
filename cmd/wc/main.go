package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

var (
	linesOpt = flag.Bool("l", false, "print the newline counts")
	wordsOpt = flag.Bool("w", false, "print the word counts")
	bytesOpt = flag.Bool("c", false, "print the byte counts")
)

func main() {
	app := cliutil.NewApp(cliutil.ProgramName(), "Print newline, word, and byte counts.")
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

	// If no options specified, default to all three
	if !*linesOpt && !*wordsOpt && !*bytesOpt {
		*linesOpt = true
		*wordsOpt = true
		*bytesOpt = true
	}

	totalLines, totalWords, totalBytes := 0, 0, 0

	for _, file := range files {
		f, err := common.OpenInput(file)
		if err != nil {
			return err
		}
		lines, words, bytes := count(f)
		f.Close()
		totalLines += lines
		totalWords += words
		totalBytes += bytes
		printCounts(lines, words, bytes, filepath.Base(file))
	}

	if len(files) > 1 {
		printCounts(totalLines, totalWords, totalBytes, "total")
	}
	return nil
}

func count(f *os.File) (int, int, int) {
	reader := bufio.NewReader(f)
	lines, words, bytes := 0, 0, 0
	inWord := false

	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		bytes++
		if b == '\n' {
			lines++
		}
		if unicode.IsSpace(rune(b)) {
			if inWord {
				inWord = false
				words++
			}
		} else {
			inWord = true
		}
	}
	if inWord {
		words++
	}
	return lines, words, bytes
}

func printCounts(lines, words, bytes int, name string) {
	var parts []string
	if *linesOpt {
		parts = append(parts, fmt.Sprintf("%8d", lines))
	}
	if *wordsOpt {
		parts = append(parts, fmt.Sprintf("%8d", words))
	}
	if *bytesOpt {
		parts = append(parts, fmt.Sprintf("%8d", bytes))
	}
	if name != "-" {
		parts = append(parts, name)
	}
	fmt.Println(strings.Join(parts, " "))
}
