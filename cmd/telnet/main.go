package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	app := cliutil.NewApp("telnet", "Simplified TCP client.")
	flag.Usage = app.Usage
	flag.Parse()

	if flag.NArg() < 2 {
		common.Warn("telnet: usage: telnet host port")
		app.Usage()
		return common.ExitFailure
	}
	addr := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		common.Warn("telnet: %v", err)
		return common.ExitFailure
	}
	defer conn.Close()
	fmt.Fprintf(os.Stderr, "Connected to %s\n", addr)
	go func() {
		ioCopy(os.Stdout, conn)
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
	return common.ExitSuccess
}

func ioCopy(dst *os.File, src net.Conn) {
	bufio.NewReader(src).WriteTo(dst)
}

func main() {
	os.Exit(run())
}
