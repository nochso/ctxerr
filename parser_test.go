package ctxerr

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/nochso/golden"
)

var update = flag.Bool("update", false, "update golden test files")

func TestParse(t *testing.T) {
	golden.TestDir(t, "test-fixtures/parse-ok", func(tc golden.Case) {
		ctx, err := Parse(tc.In.String())
		if err != nil {
			t.Errorf("expected nil error; got %#v", err.Error())
			return
		}
		if *update {
			tc.Out.Update([]byte(ctx.Error()))
		}
		tc.Diff(ctx.Error())
	})
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
	prefix := "strconv"
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Errorf("expected error to have prefix %#v, got %#v", prefix, err.Error())
	}
}
