package ctxerr

import (
	"fmt"

	"github.com/fatih/color"
)

func init() {
	color.NoColor = true
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

func ExampleCtx_WithHint() {
	ctx := New("010101102110", Point(1, 9))
	ctx.Hint = "don't worry, bender"
	fmt.Println(ctx)
	// Output:
	// 1:9:
	// 1 | 010101102110
	//   |         ^ don't worry, bender
}

func ExampleCtx_WithContext_limited() {
	in := `1st
2nd
3rd has an error
4th
5th`
	ctx := New(in, Point(3, 12)).WithContext(1)
	fmt.Println(ctx)
	// Output:
	// 3:12:
	// 2 | 2nd
	// 3 | 3rd has an error
	//   |            ^
	// 4 | 4th
}

func ExampleCtx_WithContext_limitedMultiline() {
	in := `1st
2nd
3rd has an error
4th still has an error
5th`
	ctx := New(in, Range(3, 1, 4, 22)).WithContext(1)
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

func ExampleCtx_WithContext_all() {
	in := `1st
2nd
3rd has an error
4th
5th`
	ctx := New(in, Point(3, 12)).WithContext(-1)
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

func ExampleCtx_Path() {
	ctx := New("42", Point(1, 1))
	ctx.Path = "/tmp/ctxerr/answer.txt"
	fmt.Println(ctx)
	// Output:
	// /tmp/ctxerr/answer.txt:1:1:
	// 1 | 42
	//   | ^
}

func ExampleCtx_ToError() {
	cause := fmt.Errorf("not a letter")
	err := New("ab!cd", Point(1, 3)).ToError(cause)
	fmt.Println(err)
	// Output:
	// not a letter
	// 1:3:
	// 1 | ab!cd
	//   |   ^
}
