package service

import "testing"

func TestRule180(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule180(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 180: %d", r)
		}
	}
}
func TestRule181(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule181(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 181: %d", r)
		}
	}
}
func TestRule182(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule182(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 182: %d", r)
		}
	}
}
func TestRule183(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule183(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 183: %d", r)
		}
	}
}
func TestRule184(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule184(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 184: %d", r)
		}
	}
}
func TestRule185(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule185(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 185: %d", r)
		}
	}
}
func TestRule186(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule186(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 186: %d", r)
		}
	}
}
func TestRule187(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule187(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 187: %d", r)
		}
	}
}
func TestRule188(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule188(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 188: %d", r)
		}
	}
}
func TestRule189(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule189(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 189: %d", r)
		}
	}
}
