package ctxerr

import "fmt"

type position struct {
	line, col int
}

func (p position) String() string {
	if p.col == 0 {
		return fmt.Sprintf("%d", p.line)
	}
	return fmt.Sprintf("%d:%d", p.line, p.col)
}

// Region defines a selection of runes (utf8 codepoints) in a string.
type Region struct {
	start, end position
}

// Point returns a Region pointing to a specific rune.
// Line and column are one-based.
func Point(line, col int) Region {
	return Range(line, col, line, col)
}

// Range returns a Region pointing to a range of runes.
// Line and column are one-based.
func Range(startLine, startCol, endLine, endCol int) Region {
	return Region{
		start: position{startLine, startCol},
		end:   position{endLine, endCol},
	}
}

// String representation of a region.
//
//	3       // complete 3rd line
//	1-3     // lines 1 through 3
//	1:1     // pointer
//	1:1-2   // range on single line
//	1:1-2:1 // range over multiple lines
func (r Region) String() string {
	if r.isPointer() {
		return r.start.String()
	}
	if r.isMultiLine() {
		return fmt.Sprintf("%s-%s", r.start.String(), r.end.String())
	}
	if r.end.col == 0 {
		return r.end.String()
	}
	return fmt.Sprintf("%s-%d", r.start.String(), r.end.col)
}

func (r Region) isPointer() bool {
	return r.start.line == r.end.line && r.start.col == r.end.col && r.start.col != 0
}

func (r Region) isMultiLine() bool {
	return r.start.line != r.end.line
}
