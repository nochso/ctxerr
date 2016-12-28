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
	// 1 | 00100
	//   |   ^
}

func ExampleRange() {
	fmt.Println(New("01110", Range(1, 2, 1, 4)))
	// Output:
	// 1 | 01110
	//   |  ^^^
}

func ExampleRange_multiline() {
	fmt.Println(New("00001\n11110", Range(1, 5, 2, 4)))
	// Output:
	// 1 | 00001
	//   |     ^
	// 2 | 11110
	//   | ^^^^
}

func ExampleCtx_WithHint() {
	ctx := New("010101102110", Point(1, 9)).WithHint("don't worry, bender")
	fmt.Println(ctx)
	// Output:
	// 1 | 010101102110
	//   |         ^ don't worry, bender
}
