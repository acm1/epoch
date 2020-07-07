package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	var t time.Time

	if len(os.Args) < 2 {
		t = time.Now()
	} else {
		if helpRequested(os.Args) {
			os.Exit(0)
		}

		timestamp := strings.Join(os.Args[1:], " ")

		var err error

		t, err = dateparse.ParseLocal(timestamp)
		if err != nil {
			printUsage(os.Stderr, os.Args[0])
			_, _ = fmt.Fprintf(os.Stderr, "\nfailed to parse timestamp: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Local time:   ", t)
	fmt.Println("UTC time:     ", t.UTC())
	fmt.Println("Epoch seconds:", t.Unix())
	fmt.Println("Epoch millis: ", t.UnixNano()/int64(time.Millisecond))
}

const usage =
`Usage:
    %s [timestamp]

Example:
    %[1]s 7/7/20 10:03am

Prints a time in human-readable format in local and UTC, and in Unix
Epoch format. If no arguments are given, the time printed is the
current time, otherwise the arguments are parsed as the time to print.
Parsable formats can be found at https://github.com/araddon/dateparse.
`

// helpRequested looks for "-h" or "--help" in the args slice, and if one of them is found, prints usage instructions
// to stdout and returns true
func helpRequested(args []string) bool {
	for _, arg := range args[1:] {
		if arg == "-h" || arg == "--help" {
			printUsage(os.Stdout, args[0])
			return true
		}
	}

	return false
}

func printUsage(w io.Writer, cmd string) {
	_, _ = fmt.Fprintf(w, usage, cmd)
}
