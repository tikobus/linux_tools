package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		listen = flag.Bool("l", false, "listen mode")
		udp    = flag.Bool("u", false, "use UDP")
		wait   = flag.Duration("w", 3*time.Second, "timeout for connects")
		zeroIO = flag.Bool("z", false, "zero-I/O mode (scanning)")
	)
	app := cliutil.NewApp("nc", "Netcat: TCP/UDP swiss army knife.")
	flag.Usage = app.Usage
	flag.Parse()

	proto := "tcp"
	if *udp {
		proto = "udp"
	}

	if *listen {
		port := flag.Arg(0)
		if port == "" {
			common.Warn("nc: missing port")
			return common.ExitFailure
		}
		addr := ":" + port
		l, err := net.Listen(proto, addr)
		if err != nil {
			common.Warn("nc: %v", err)
			return common.ExitFailure
		}
		defer l.Close()
		fmt.Fprintf(os.Stderr, "Listening on %s %s\n", proto, addr)
		conn, err := l.Accept()
		if err != nil {
			common.Warn("nc: %v", err)
			return common.ExitFailure
		}
		defer conn.Close()
		io.Copy(os.Stdout, conn)
		return common.ExitSuccess
	}

	if flag.NArg() < 2 {
		common.Warn("nc: usage: nc host port")
		app.Usage()
		return common.ExitFailure
	}
	host := flag.Arg(0)
	port := flag.Arg(1)
	addr := net.JoinHostPort(host, port)

	if *zeroIO {
		conn, err := net.DialTimeout(proto, addr, *wait)
		if err != nil {
			fmt.Printf("nc: connect to %s failed: %v\n", addr, err)
			return common.ExitFailure
		}
		conn.Close()
		fmt.Printf("Connection to %s succeeded.\n", addr)
		return common.ExitSuccess
	}

	conn, err := net.Dial(proto, addr)
	if err != nil {
		common.Warn("nc: %v", err)
		return common.ExitFailure
	}
	defer conn.Close()
	go func() { io.Copy(conn, os.Stdin) }()
	io.Copy(os.Stdout, conn)
	return common.ExitSuccess
}

func main() {
	os.Exit(run())
}
