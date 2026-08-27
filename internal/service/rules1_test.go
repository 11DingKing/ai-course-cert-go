package service

import "testing"

func TestRule20(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule20(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 20: %d", r)
		}
	}
}
func TestRule21(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule21(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 21: %d", r)
		}
	}
}
func TestRule22(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule22(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 22: %d", r)
		}
	}
}
func TestRule23(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule23(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 23: %d", r)
		}
	}
}
func TestRule24(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule24(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 24: %d", r)
		}
	}
}
func TestRule25(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule25(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 25: %d", r)
		}
	}
}
func TestRule26(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule26(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 26: %d", r)
		}
	}
}
func TestRule27(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule27(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 27: %d", r)
		}
	}
}
func TestRule28(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule28(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 28: %d", r)
		}
	}
}
func TestRule29(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule29(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 29: %d", r)
		}
	}
}
