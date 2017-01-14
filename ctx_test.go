package ctxerr

import (
	"fmt"
	"os"
	"testing"

	"github.com/fatih/color"
)

func init() {
	color.NoColor = true
}

func ExampleNew() {
	fmt.Println(New("source code containing error", Range(1, 24, 1, 28)))
	// Output:
	// 1:24-28:
	// 1 | source code containing error
	//   |                        ^^^^^
}

func ExamplePoint() {
	fmt.Println(New("00100", Point(1, 3)))
	// Output:
	// 1:3:
	// 1 | 00100
	//   |   ^
}

func ExampleRange() {
	fmt.Println(New("01110", Range(1, 2, 1, 4)))
	// Output:
	// 1:2-4:
	// 1 | 01110
	//   |  ^^^
}

func ExampleRange_multiline() {
	fmt.Println(New("00001\n11110", Range(1, 5, 2, 4)))
	// Output:
	// 1:5-2:4:
	// 1 | 00001
	//   |     ^
	// 2 | 11110
	//   | ^^^^
}

func ExampleCtx_Error_hint() {
	ctx := New("010101102110", Point(1, 9))
	ctx.Hint = "don't worry, bender"
	fmt.Println(ctx)
	// Output:
	// 1:9:
	// 1 | 010101102110
	//   |         ^ don't worry, bender
}

func ExampleCtx_Error_contextLimited() {
	in := `1st
2nd
3rd has an error
4th
5th`
	ctx := New(in, Point(3, 12))
	ctx.Context = 1
	fmt.Println(ctx)
	// Output:
	// 3:12:
	// 2 | 2nd
	// 3 | 3rd has an error
	//   |            ^
	// 4 | 4th
}

func ExampleCtx_Error_contextLimitedMultiline() {
	in := `1st
2nd
3rd has an error
4th still has an error
5th`
	ctx := New(in, Range(3, 1, 4, 22))
	ctx.Context = 1
	fmt.Println(ctx)
	// Output:
	// 3:1-4:22:
	// 2 | 2nd
	// 3 | 3rd has an error
	//   | ^^^^^^^^^^^^^^^^
	// 4 | 4th still has an error
	//   | ^^^^^^^^^^^^^^^^^^^^^^
	// 5 | 5th
}

func ExampleCtx_Error_contextAll() {
	in := `1st
2nd
3rd has an error
4th
5th`
	ctx := New(in, Point(3, 12))
	ctx.Context = -1
	fmt.Println(ctx)
	// Output:
	// 3:12:
	// 1 | 1st
	// 2 | 2nd
	// 3 | 3rd has an error
	//   |            ^
	// 4 | 4th
	// 5 | 5th
}

func ExampleCtx_Error_path() {
	ctx := New("42", Point(1, 1))
	ctx.Path = "/tmp/ctxerr/answer.txt"
	fmt.Println(ctx)
	// Output:
	// /tmp/ctxerr/answer.txt:1:1:
	// 1 | 42
	//   | ^
}

func ExampleCtx_Error() {
	err := New("ab!cd", Point(1, 3))
	err.Err = fmt.Errorf("not a letter")
	fmt.Println(err)
	// Output:
	// not a letter
	// 1:3:
	// 1 | ab!cd
	//   |   ^
}

func ExampleCtx_Error_tabwidth() {
	in := "\tfoo\tbar"
	ctx := New(in, Point(1, 5))
	fmt.Println(ctx)
	// Output:
	// 1:5:
	// 1 |     foo    bar
	//   |        ^^^^
}

func ExampleNewFromPath() {
	cerr, err := NewFromPath("LICENSE", Range(1, 1, 1, 3))
	if err != nil {
		fmt.Println(err)
		return
	}
	cerr.Context = 0
	fmt.Println(cerr)
	// Output:
	// LICENSE:1:1-3:
	// 1 | MIT License
	//   | ^^^
}

func TestNewFromPath_NoSuchFile(t *testing.T) {
	_, err := NewFromPath("foo.txt", Point(1, 1))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected %#v, got %#v", os.ErrNotExist.Error(), err.Error())
	}
}

func testCmp(t *testing.T, exp, act string) {
	if exp != act {
		t.Errorf("expected %#v, got %#v", exp, act)
	}
}

func TestCtx_Error_contextBoundaries(t *testing.T) {
	ctx := New("x", Point(1, 1))
	ctx.Context = 42
	exp := `1:1:
1 | x
  | ^
`
	testCmp(t, exp, ctx.Error())
}

func TestCtx_Error_nilColumn(t *testing.T) {
	ctx := New("foo", Point(1, 0))
	exp := `1:
1 | foo
  | ^^^
`
	testCmp(t, exp, ctx.Error())
}
