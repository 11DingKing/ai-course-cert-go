package service

import "testing"

func TestRule80(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule80(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 80: %d", r)
		}
	}
}
func TestRule81(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule81(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 81: %d", r)
		}
	}
}
func TestRule82(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule82(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 82: %d", r)
		}
	}
}
func TestRule83(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule83(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 83: %d", r)
		}
	}
}
func TestRule84(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule84(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 84: %d", r)
		}
	}
}
func TestRule85(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule85(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 85: %d", r)
		}
	}
}
func TestRule86(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule86(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 86: %d", r)
		}
	}
}
func TestRule87(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule87(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 87: %d", r)
		}
	}
}
func TestRule88(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule88(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 88: %d", r)
		}
	}
}
func TestRule89(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule89(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 89: %d", r)
		}
	}
}
