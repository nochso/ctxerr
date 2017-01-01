package ctxerr

import "fmt"

// Error points at the cause of any kind of parsing error.
type Error struct {
	Ctx   Ctx
	Inner error
}

// NewError creates a new error with additional context.
func NewError(err error, ctx Ctx) Error {
	return Error{ctx, err}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s\n%s", e.Inner, e.Ctx)
}
