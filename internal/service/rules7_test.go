package service

import "testing"

func TestRule140(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule140(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 140: %d", r)
		}
	}
}
func TestRule141(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule141(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 141: %d", r)
		}
	}
}
func TestRule142(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule142(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 142: %d", r)
		}
	}
}
func TestRule143(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule143(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 143: %d", r)
		}
	}
}
func TestRule144(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule144(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 144: %d", r)
		}
	}
}
func TestRule145(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule145(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 145: %d", r)
		}
	}
}
func TestRule146(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule146(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 146: %d", r)
		}
	}
}
func TestRule147(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule147(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 147: %d", r)
		}
	}
}
func TestRule148(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule148(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 148: %d", r)
		}
	}
}
func TestRule149(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule149(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 149: %d", r)
		}
	}
}
