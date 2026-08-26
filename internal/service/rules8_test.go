package service

import "testing"

func TestRule160(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule160(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 160: %d", r)
		}
	}
}
func TestRule161(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule161(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 161: %d", r)
		}
	}
}
func TestRule162(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule162(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 162: %d", r)
		}
	}
}
func TestRule163(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule163(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 163: %d", r)
		}
	}
}
func TestRule164(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule164(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 164: %d", r)
		}
	}
}
func TestRule165(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule165(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 165: %d", r)
		}
	}
}
func TestRule166(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule166(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 166: %d", r)
		}
	}
}
func TestRule167(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule167(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 167: %d", r)
		}
	}
}
func TestRule168(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule168(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 168: %d", r)
		}
	}
}
func TestRule169(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule169(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 169: %d", r)
		}
	}
}
