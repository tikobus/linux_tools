package common

import (
	"bufio"
	"io"
	"os"
)

// CopyStream copies all data from src to dst.
func CopyStream(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

// ReadLines reads all lines from r.
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// OpenInput opens filename or stdin when filename == "-".
func OpenInput(filename string) (*os.File, error) {
	if filename == "-" {
		return os.Stdin, nil
	}
	return os.Open(filename)
}

// CreateOutput creates filename or returns stdout when filename == "-".
func CreateOutput(filename string) (*os.File, error) {
	if filename == "-" {
		return os.Stdout, nil
	}
	return os.Create(filename)
}
