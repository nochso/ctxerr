package ctxerr

import (
	"os"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	ctx, err := Parse("parser_test.go:1:8: Was habbe se denn? - Abitur!")
	if err != nil {
		t.Error(err)
	}
	if ctx == nil {
		t.Error("expected non-nil Ctx")
	}
	expected := `Was habbe se denn? - Abitur!
parser_test.go:1:8:
1 | package ctxerr
  |        ^
`
	if ctx.Error() != expected {
		t.Errorf("expected:\n%#v\ngot:\n%#v", expected, ctx.Error())
	}
}

func TestParse_NoMatch(t *testing.T) {
	input := []string{
		"",
		" ",
		":",
		"a:",
		"a:5",
	}
	for _, in := range input {
		_, err := Parse(in)
		if err != ErrNoMatch {
			t.Errorf("expected %#v, got %#v", ErrNoMatch.Error(), err.Error())
		}
	}
}

func TestParse_NoSuchFile(t *testing.T) {
	_, err := Parse("foo.txt:1:")
	if !os.IsNotExist(err) {
		t.Errorf("expected %#v, got %#v", os.ErrNotExist.Error(), err.Error())
	}
}

func TestParse_InvalidInt(t *testing.T) {
	_, err := Parse("a:9999999999999999999999:")
	if err == nil {
		t.Error("expected error, got nil")
	}
	prefix := "error parsing line number"
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Errorf("expected error to have prefix %#v, got %#v", prefix, err.Error())
	}
}
