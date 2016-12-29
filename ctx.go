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

var (
	// DefaultGutter separates line numbers from content.
	DefaultGutter = " | "
	// DefaultPointer is the rune that points up at a region.
	DefaultPointer = '^'
)

// Ctx points to runes in (multiline) strings.
type Ctx struct {
	lines   []string
	context int
	Region
	path string
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

// WithContext returns a Ctx with a maximum amount of context lines.
//
//	 0: no lines of context.
//	-1: all lines, the full input string
//	 3: limited context of 3 lines
func (c Ctx) WithContext(context int) Ctx {
	c.context = context
	return c
}

// WithPath returns a Ctx with the given path added to the region string.
func (c Ctx) WithPath(path string) Ctx {
	c.path = path
	return c
}

func split(s string, r Region) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	l := []string{}
	for sc.Scan() {
		l = append(l, sc.Text())
	}
	return l
}

// ToError wraps err with this context.
func (c Ctx) ToError(err error) CtxErr {
	return CtxErr{
		ctx: c,
		err: err,
	}
}

func (c Ctx) String() string {
	buf := &bytes.Buffer{}
	if c.path != "" {
		fmt.Fprintf(buf, "%s:", c.path)
	}
	fmt.Fprintf(buf, "%s:\n", c.Region)
	start, end := c.lineIndex()
	// length of highest line number
	linePosMaxLen := posLen(end)
	for i, line := range c.lines[start:end] {
		linePos := start + i + 1
		// write line no. gutter and actual line
		c.writeLineGutter(buf, linePos, linePosMaxLen)
		buf.WriteString(line)
		buf.WriteByte('\n')
		if linePos < c.start.line || linePos > c.end.line {
			// this was just context, don't point at it
			continue
		}
		// this line is being pointed at
		c.writeLineGutter(buf, 0, linePosMaxLen)
		buf.WriteString(strings.Repeat(" ", c.getPad(linePos)))
		buf.WriteString(color.RedString("%s", strings.Repeat(string(DefaultPointer), c.getDots(linePos, line))))
		if c.hint != "" && c.start.line == linePos {
			fmt.Fprintf(buf, " %s", c.hint)
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// start and end index of Ctx.lines including lines of context.
func (c Ctx) lineIndex() (start, end int) {
	if c.context < 0 {
		return 0, len(c.lines)
	}
	start = c.start.line - c.context - 1
	if start < 0 {
		start = 0
	}
	end = c.end.line + c.context
	if end > len(c.lines) {
		end = len(c.lines)
	}
	return
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
	buf.WriteString(DefaultGutter)
}
