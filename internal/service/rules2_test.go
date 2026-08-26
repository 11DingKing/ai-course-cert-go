package service

import "testing"

func TestRule40(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule40(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 40: %d", r)
		}
	}
}
func TestRule41(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule41(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 41: %d", r)
		}
	}
}
func TestRule42(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule42(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 42: %d", r)
		}
	}
}
func TestRule43(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule43(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 43: %d", r)
		}
	}
}
func TestRule44(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule44(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 44: %d", r)
		}
	}
}
func TestRule45(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule45(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 45: %d", r)
		}
	}
}
func TestRule46(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule46(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 46: %d", r)
		}
	}
}
func TestRule47(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule47(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 47: %d", r)
		}
	}
}
func TestRule48(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule48(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 48: %d", r)
		}
	}
}
func TestRule49(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule49(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 49: %d", r)
		}
	}
}
