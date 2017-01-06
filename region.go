package ctxerr

import "fmt"

type Position struct {
	Line, Col int
}

func (p Position) String() string {
	if p.Col == 0 {
		return fmt.Sprintf("%d", p.Line)
	}
	return fmt.Sprintf("%d:%d", p.Line, p.Col)
}

// Region defines a selection of runes (utf8 codepoints) in a string.
type Region struct {
	start, end Position
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
		start: Position{startLine, startCol},
		end:   Position{endLine, endCol},
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
	if r.end.Col == 0 {
		return r.end.String()
	}
	return fmt.Sprintf("%s-%d", r.start.String(), r.end.Col)
}

func (r Region) isPointer() bool {
	return r.start.Line == r.end.Line && r.start.Col == r.end.Col && r.start.Col != 0
}

func (r Region) isMultiLine() bool {
	return r.start.Line != r.end.Line
}
