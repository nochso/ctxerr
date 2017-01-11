package ctxerr

import "fmt"

// Position of a single point in a multiline string.
type Position struct {
	Line, Col int
}

// String representation of a Position.
//
//	1
//	1:2
func (p Position) String() string {
	if p.Col == 0 {
		return fmt.Sprintf("%d", p.Line)
	}
	return fmt.Sprintf("%d:%d", p.Line, p.Col)
}

// Region defines a selection of runes (utf8 codepoints) in a string.
type Region struct {
	Start, End Position
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
		Start: Position{startLine, startCol},
		End:   Position{endLine, endCol},
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
		return r.Start.String()
	}
	if r.isMultiLine() {
		return fmt.Sprintf("%s-%s", r.Start.String(), r.End.String())
	}
	if r.End.Col == 0 {
		return r.End.String()
	}
	return fmt.Sprintf("%s-%d", r.Start.String(), r.End.Col)
}

func (r Region) isPointer() bool {
	return r.Start.Line == r.End.Line && r.Start.Col == r.End.Col && r.Start.Col != 0
}

func (r Region) isMultiLine() bool {
	return r.Start.Line != r.End.Line
}
