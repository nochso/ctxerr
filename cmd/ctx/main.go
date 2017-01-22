package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nochso/ctxerr"
)

var (
	// Version as specified by ldflags on build or this default.
	Version = "1.0.0"
	// BuildDate is optional and can be supplied by ldflags.
	BuildDate = ""
)

var (
	context     = flag.Int("context", ctxerr.DefaultContext, "print `NUM` lines of context surrounding an error. negative for all, positive for limited and 0 for none (default)")
	pessimistic = flag.Bool("pessimistic", false, "print only matching errors, ignore everything else")
)

func usage() {
	bd := BuildDate
	if bd != "" {
		bd = fmt.Sprintf(" (built %s)", bd)
	}
	format := `ctx %s%s
Pretty prints parser errors from stdin.

Possible input:
  path/file.ext:1:5: some error on line 1, column 5
  file.ext:123: column is optional and so is the message:
  file.ext:1

Usage:
  ctx < log.txt
  gometalinter . | ctx

Flags:
`
	fmt.Printf(format, Version, bd)
	flag.PrintDefaults()
}

func main() {
	flag.BoolVar(&ctxerr.NoColor, "no-color", false, "disable any color output")
	flag.Usage = usage
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		cerr, err := ctxerr.Parse(sc.Text())
		if err != nil {
			if err != ctxerr.ErrNoMatch {
				fmt.Println(err)
			}
			if !*pessimistic {
				fmt.Println(sc.Text())
			}
			continue
		}
		cerr.Context = *context
		fmt.Println(cerr)
	}

	if sc.Err() != nil {
		fmt.Fprint(os.Stderr, sc.Err())
		os.Exit(1)
	}
}
