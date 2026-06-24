package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/user/linux_coreutils/pkg/cliutil"
	"github.com/user/linux_coreutils/pkg/common"
)

func run() int {
	var (
		showYear = flag.Bool("y", false, "show whole year")
	)
	app := cliutil.NewApp("cal", "Display a calendar.")
	flag.Usage = app.Usage
	flag.Parse()

	now := time.Now()
	month := now.Month()
	year := now.Year()

	switch flag.NArg() {
	case 0:
		// default to current month
	case 1:
		// year only
		y, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			common.Warn("cal: invalid year %q", flag.Arg(0))
			return common.ExitFailure
		}
		year = y
		*showYear = true
	case 2:
		// month and year
		m, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			common.Warn("cal: invalid month %q", flag.Arg(0))
			return common.ExitFailure
		}
		if m < 1 || m > 12 {
			common.Warn("cal: month %d is out of range", m)
			return common.ExitFailure
		}
		month = time.Month(m)
		y, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			common.Warn("cal: invalid year %q", flag.Arg(1))
			return common.ExitFailure
		}
		year = y
	default:
		common.Warn("cal: too many arguments")
		app.Usage()
		return common.ExitFailure
	}

	if *showYear {
		printYear(year)
	} else {
		printMonth(year, month)
	}
	return common.ExitSuccess
}

func printMonth(year int, month time.Month) {
	fmt.Printf("     %s %d\n", month.String(), year)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	first := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := first.AddDate(0, 1, -1).Day()

	startPad := int(first.Weekday())
	for i := 0; i < startPad; i++ {
		fmt.Print("   ")
	}
	for d := 1; d <= lastDay; d++ {
		fmt.Printf("%2d ", d)
		if (startPad+d)%7 == 0 {
			fmt.Println()
		}
	}
	if (startPad+lastDay)%7 != 0 {
		fmt.Println()
	}
}

func printYear(year int) {
	for m := time.January; m <= time.December; m++ {
		printMonth(year, m)
		if m != time.December {
			fmt.Println()
		}
	}
}

func main() {
	os.Exit(run())
}
