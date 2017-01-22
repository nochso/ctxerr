package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

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
Pretty prints parser errors from stdin or any given command.

Possible input:
  path/file.ext:1:5: some error on line 1, column 5
  file.ext:123: column is optional and so is the message:
  file.ext:1
  parser_test.go[1, 15]: Missing semicolon

Usage:
  ctx < log.txt         # Existing file
  gometalinter . | ctx  # Stdin
  ctx go test           # Separate stdin and stderr of given command
  go test 2>&1 | ctx    # Combine stdin and stderr

Flags:
`
	fmt.Printf(format, Version, bd)
	flag.PrintDefaults()
}

func main() {
	flag.BoolVar(&ctxerr.NoColor, "no-color", false, "disable any color output")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		err := scan(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	scanCmd()
}

func scan(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		cerr, err := ctxerr.Parse(sc.Text())
		if err != nil {
			if err != ctxerr.ErrNoMatch {
				fmt.Fprintln(w, err)
			}
			if !*pessimistic {
				fmt.Fprintln(w, sc.Text())
			}
			continue
		}
		cerr.Context = *context
		fmt.Fprintln(w, cerr)
	}
	return sc.Err()
}

func scanCmd() {
	cmd := exec.Command(flag.Arg(0), flag.Args()[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	errs := make(chan error)
	go func() { errs <- scan(stdout, os.Stdout) }()
	go func() { errs <- scan(stderr, os.Stderr) }()
	go func() {
		wg.Wait()
		close(errs)
	}()
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	for err := range errs {
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}
}
