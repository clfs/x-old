package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	formatFlag  = flag.String("f", "small", "")
	bracketFlag = flag.Bool("b", false, "")

	styles = map[string]func(time.Time) string{
		"small":  func(t time.Time) string { return t.Format("3:04pm") },
		"medium": func(t time.Time) string { return t.Format("15:04:05.000") },
		"large":  func(t time.Time) string { return t.Format("2006-01-02T15:04:05.999999999Z07:00") },
		"unix":   func(t time.Time) string { return strconv.FormatInt(t.Unix(), 10) },
	}
)

const usage = `Usage:
    ... | stamp [-f FORMAT] [-b]

Options:
    -b          Place the timestamp in brackets.
    -f FORMAT   Choose a timestamp format (optional; default "medium").
    -h          Show this help message.

Formats:
    small       10:33am
    medium      10:33:36.042
    large       2020-11-25T10:33:55.740453-08:00
    unix        1606329242

Example:
    $ tail -f important.log | stamp -f small -b
    [10:26am] An event.
    [10:35am] Another event, sometime later.`

func main() {
	flag.Usage = func() { fmt.Fprintln(os.Stderr, usage) }
	flag.Parse()

	lineFormat := "%s %s"
	if *bracketFlag {
		lineFormat = "[%s] %s"
	}

	style, ok := styles[*formatFlag]
	if !ok {
		fmt.Fprintln(os.Stderr, "unknown time format:", *formatFlag)
		os.Exit(1)
	}

	r := bufio.NewReader(os.Stdin)
	out := os.Stdout

	for {
		// Note that ReadString has no read limit, so very long lines
		// will probably cause memory problems.
		switch line, err := r.ReadString('\n'); err {

		case nil:
			s := fmt.Sprintf(lineFormat, style(time.Now()), line)
			if _, err = out.WriteString(s); err != nil {
				fmt.Fprintln(os.Stderr, "stdout write failed:", err)
				os.Exit(1)
			}

		case io.EOF:
			os.Exit(0)

		default:
			fmt.Fprintln(os.Stderr, "unexpected error:", err)
			os.Exit(1)
		}
	}
}
