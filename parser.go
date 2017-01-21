package ctxerr

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	// ErrNoMatch is returned by Parse when no file path or position can be found.
	ErrNoMatch       = fmt.Errorf("line does not match position pattern: %s", matchRegionColon.String())
	matchRegionColon = regexp.MustCompile(`^(.+?):(\d+):(\d+)?:?\s*(.*)$`)
	matchRegionNpm   = regexp.MustCompile(`^(.+?)\[(\d+), (\d+)\]: (.*)$`)
)

// Parse a single line with at least a file path and position.
// Returns error ErrNoMatch when this does not appear to be an error line.
// Otherwise will return parsing errors or nil on success.
//
//	foo.go:15:1: message
//	foo.go[1, 2]: message
func Parse(line string) (*Ctx, error) {
	ctx, err := parseRegionColon(line)
	if err != ErrNoMatch {
		return ctx, err
	}
	ctx, err = parseRegionNpm(line)
	return ctx, err
}

func parseRegionColon(line string) (*Ctx, error) {
	m := matchRegionColon.FindStringSubmatch(line)
	if len(m) == 0 {
		return nil, ErrNoMatch
	}

	// require a line number, column may be zero
	lineNo, err := strconv.Atoi(m[2])
	col, _ := strconv.Atoi(m[3])
	if err != nil {
		return nil, err
	}
	ctx, err := NewFromPath(m[1], Point(lineNo, col))
	ctx.Err = errors.New(m[4])
	return &ctx, err
}

func parseRegionNpm(line string) (*Ctx, error) {
	m := matchRegionNpm.FindStringSubmatch(line)
	if len(m) == 0 {
		return nil, ErrNoMatch
	}

	// require a line number, column may be zero
	lineNo, err := strconv.Atoi(m[2])
	col, _ := strconv.Atoi(m[3])
	if err != nil {
		return nil, err
	}
	ctx, err := NewFromPath(m[1], Point(lineNo, col))
	ctx.Err = errors.New(m[4])
	return &ctx, err
}
