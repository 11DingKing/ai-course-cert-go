package pagination

import "testing"

func TestNormalize(t *testing.T) {
	for _, c := range []struct{ l, o, el, eo int }{{0, 0, 20, 0}, {-1, -2, 20, 0}, {200, 3, 20, 3}, {10, 5, 10, 5}} {
		p := Normalize(c.l, c.o)
		if p.Limit != c.el || p.Offset != c.eo {
			t.Fatalf("%+v", p)
		}
	}
}
