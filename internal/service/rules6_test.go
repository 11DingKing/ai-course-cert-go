package service

import "testing"

func TestRule120(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule120(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 120: %d", r)
		}
	}
}
func TestRule121(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule121(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 121: %d", r)
		}
	}
}
func TestRule122(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule122(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 122: %d", r)
		}
	}
}
func TestRule123(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule123(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 123: %d", r)
		}
	}
}
func TestRule124(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule124(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 124: %d", r)
		}
	}
}
func TestRule125(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule125(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 125: %d", r)
		}
	}
}
func TestRule126(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule126(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 126: %d", r)
		}
	}
}
func TestRule127(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule127(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 127: %d", r)
		}
	}
}
func TestRule128(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule128(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 128: %d", r)
		}
	}
}
func TestRule129(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule129(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 129: %d", r)
		}
	}
}
