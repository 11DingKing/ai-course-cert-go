package service

import "testing"

func TestRule220(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule220(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 220: %d", r)
		}
	}
}
func TestRule221(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule221(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 221: %d", r)
		}
	}
}
func TestRule222(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule222(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 222: %d", r)
		}
	}
}
func TestRule223(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule223(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 223: %d", r)
		}
	}
}
func TestRule224(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule224(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 224: %d", r)
		}
	}
}
func TestRule225(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule225(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 225: %d", r)
		}
	}
}
func TestRule226(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule226(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 226: %d", r)
		}
	}
}
func TestRule227(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule227(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 227: %d", r)
		}
	}
}
func TestRule228(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule228(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 228: %d", r)
		}
	}
}
func TestRule229(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule229(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 229: %d", r)
		}
	}
}
