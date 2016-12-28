package ctxerr

import "testing"

func TestRegion(t *testing.T) {
	tests := []struct {
		in  Region
		exp string
	}{
		{Point(1, 1), "1:1"},
		{Point(1, 5), "1:5"},
		{Range(1, 1, 1, 5), "1:1-5"},
		{Range(1, 1, 5, 1), "1:1-5:1"},
	}
	for _, test := range tests {
		act := test.in.String()
		if test.exp != act {
			t.Errorf("expected %#v; got %#v", test.exp, act)
		}
	}
}
