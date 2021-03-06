package ctxerr

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
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
	// DefaultContext is the default amount of context lines surrounding an error.
	// It is used by New* functions.
	DefaultContext = 0
	// NoColor disables color output.
	NoColor bool
)

var (
	red = color.New(color.FgRed)
)

// Ctx points to runes in (multiline) strings.
type Ctx struct {
	Lines []string
	// Context is the maximum amount of context lines surrounding the marked region.
	//
	//	 0: no lines of context.
	//	-1: all lines, the full input string
	//	 3: limited context of 3 lines
	Context int
	Region
	// Path to the source of the context (optional).
	Path string
	// Hint that is displayed near the region markers (optional).
	Hint string
	// Err is the error that occurred at this region (optional).
	Err error
}

// New returns a new Ctx pointing to a region in an input string.
//
// Use functions Point and Range to create a Region.
func New(input string, region Region) Ctx {
	return Ctx{
		Lines:   split(input, region),
		Region:  region,
		Context: DefaultContext,
	}
}

// NewFromPath returns a new Ctx pointing to a region in the given file.
// Returns an error when the file does not exist or could not be read.
func NewFromPath(path string, region Region) (Ctx, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Ctx{}, err
	}
	c := New(string(b), region)
	c.Path = path
	return c, nil
}

func split(s string, r Region) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	l := []string{}
	for sc.Scan() {
		l = append(l, sc.Text())
	}
	return l
}

// Error formats a summary of this context error.
func (c Ctx) Error() string {
	buf := &bytes.Buffer{}
	if c.Err != nil {
		fmt.Fprintf(buf, "%s\n", c.Err)
	}
	if c.Path != "" {
		fmt.Fprintf(buf, "%s:", c.Path)
	}
	fmt.Fprintf(buf, "%s:\n", c.Region)
	start, end := c.lineIndex()
	// length of highest line number
	linePosMaxLen := posLen(end)
	for i, line := range c.Lines[start:end] {
		linePos := start + i + 1
		// write line no. gutter and actual line
		c.writeLineGutter(buf, linePos, linePosMaxLen)
		buf.WriteString(strings.Replace(line, "\t", "    ", -1))
		buf.WriteByte('\n')
		if linePos < c.Start.Line || linePos > c.End.Line {
			// this was just context, don't point at it
			continue
		}
		// this line is being pointed at
		c.writeLineGutter(buf, 0, linePosMaxLen)
		buf.WriteString(strings.Repeat(" ", c.getPad(linePos)))
		buf.WriteString(paint(red, "%s", strings.Repeat(string(DefaultPointer), c.getDots(linePos, line))))
		if c.Hint != "" && c.Start.Line == linePos {
			fmt.Fprintf(buf, " %s", c.Hint)
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// ErrorLine returns a single line error message.
func (c Ctx) ErrorLine() string {
	s := c.Path
	if s != "" {
		s += ":"
	}
	s += c.Region.String() + ": " + c.Err.Error()
	return s
}

func paint(c *color.Color, format string, a ...interface{}) string {
	if NoColor {
		return fmt.Sprintf(format, a...)
	}
	return c.SprintfFunc()(format, a...)
}

// start and end index of Ctx.Lines including lines of context.
func (c Ctx) lineIndex() (start, end int) {
	if c.Context < 0 {
		return 0, len(c.Lines)
	}
	start = c.Start.Line - c.Context - 1
	if start < 0 {
		start = 0
	}
	end = c.End.Line + c.Context
	if end > len(c.Lines) {
		end = len(c.Lines)
	}
	return
}

func posLen(i int) int {
	return len(strconv.Itoa(i))
}

func (c Ctx) getDots(pos int, line string) int {
	start := 0
	end := len(line)
	if c.Start.Line == pos && c.Start.Col != 0 {
		start = c.Start.Col - 1
	}
	if c.End.Line == pos && c.End.Col != 0 {
		end = c.End.Col
	}

	// How far do we go "over the line"?
	// Remember the extra and limit the end to avoid index errors.
	extra := 0
	if end > len(line) {
		extra = end - len(line)
		end = len(line)
	}
	s := line[start:end]
	// add extra dots over the line
	return extra + utf8.RuneCountInString(s) + strings.Count(s, "\t")*3
}

func (c Ctx) getPad(pos int) int {
	if c.isMultiLine() && c.Start.Line != pos {
		return 0
	}
	if (c.Start.Line == pos && c.Start.Col == 0) || (c.End.Line == pos && c.End.Col == 0) {
		return 0
	}
	pad := c.Start.Col - 1
	pad += strings.Count(c.Lines[c.Start.Line-1][:pad], "\t") * 3
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
