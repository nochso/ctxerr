package ctxerr

import (
	"fmt"
	"testing"
)

func ExampleError_Error() {
	cause := fmt.Errorf("foo must not be negative")
	ctx := New("foo = -1", Point(1, 7))
	err := NewError(cause, ctx)
	fmt.Println(err)
	// Output:
	// foo must not be negative
	// 1:7:
	// 1 | foo = -1
	//   |       ^
}

func TestError_Inner(t *testing.T) {
	cause := fmt.Errorf("foo must not be negative")
	ctx := New("foo = -1", Point(1, 7))
	err := NewError(cause, ctx)
	if err.Inner() != cause {
		t.Fatalf("expected %s; got %s", cause, err.Inner())
	}
}
