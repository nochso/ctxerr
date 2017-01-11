package ctxerr

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var (
	// ErrNoMatch is returned by Parse when no file path or position can be found.
	ErrNoMatch      = fmt.Errorf("line does not match position pattern: %s", positionPattern.String())
	positionPattern = regexp.MustCompile(`^(.+?):(\d+):(\d+)?:?\s*(.*)$`)
)

// Parse a single line with at least a file path and position.
// Returns error ErrNoMatch when this does not appear to be an error line.
// Otherwise will return parsing errors or nil on success.
//
//	ctx, err := ctxerr.Parse(`foo.go:15:1:`)
func Parse(line string) (*Ctx, error) {
	m := positionPattern.FindStringSubmatch(line)
	if len(m) == 0 {
		return nil, ErrNoMatch
	}

	// require a line number, column may be zero
	lineNo, err := strconv.Atoi(m[2])
	col, _ := strconv.Atoi(m[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing line number %#v: %s", m[2], err)
	}
	region := Point(lineNo, col)

	fpath := m[1]
	fbytes, err := ioutil.ReadFile(fpath)
	ctx := New(string(fbytes), region)
	ctx.Path = fpath
	ctx.Err = fmt.Errorf("%s", errors.New(m[4]))
	return &ctx, err
}
