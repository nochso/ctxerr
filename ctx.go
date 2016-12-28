package ctxerr

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

// Ctx points to runes in (multiline) strings.
type Ctx struct {
	lines []string
	Region
	hint string
}

// New returns a new Ctx pointing to a region in input.
//
// See functions Point and Range.
func New(input string, region Region) Ctx {
	return Ctx{
		lines:  split(input, region),
		Region: region,
	}
}

// WithHint returns a Ctx with a text hint that is displayed near the region markers.
func (c Ctx) WithHint(hint string) Ctx {
	c.hint = hint
	return c
}

func split(s string, r Region) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	l := make([]string, 0, r.end.line-r.start.line+1)
	line := 1
	for sc.Scan() {
		if line >= r.start.line && line <= r.end.line {
			l = append(l, sc.Text())
		}
		line++
		if line > r.end.line {
			break
		}
	}
	return l
}

func (c Ctx) String() string {
	buf := &bytes.Buffer{}
	linePosMaxLen := posLen(c.end.line)
	for i, line := range c.lines {
		linePos := c.start.line + i
		c.writeLineGutter(buf, linePos, linePosMaxLen)
		buf.WriteString(line)
		buf.WriteByte('\n')

		c.writeLineGutter(buf, 0, linePosMaxLen)
		buf.WriteString(strings.Repeat(" ", c.getPad(linePos)))
		buf.WriteString(color.RedString("%s", strings.Repeat("^", c.getDots(linePos, line))))
		if c.hint != "" && c.start.line == linePos {
			fmt.Fprintf(buf, " %s", c.hint)
		}
		if linePos < c.end.line {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

func posLen(i int) int {
	return len(strconv.Itoa(i))
}

func (c Ctx) getDots(pos int, line string) int {
	if c.isPointer() {
		return 1
	}
	if !c.isMultiLine() {
		return c.end.col - c.start.col + 1
	}
	if c.start.line == pos {
		return utf8.RuneCountInString(line) - c.start.col + 1
	}
	if c.end.line == pos {
		return c.end.col
	}
	return utf8.RuneCountInString(line)
}

func (c Ctx) getPad(pos int) int {
	pad := c.start.col - 1
	if c.isMultiLine() && c.start.line != pos {
		pad = 0
	}
	return pad
}

func (c Ctx) writeLineGutter(buf *bytes.Buffer, line, maxLen int) {
	pad := maxLen  // assume 0, meaning no line info
	if line != 0 { // otherwise exclude line no. from padding
		pad -= posLen(line)
	}
	buf.WriteString(strings.Repeat(" ", pad))
	if line != 0 {
		buf.WriteString(strconv.Itoa(line))
	}
	buf.WriteString(" | ")
}
