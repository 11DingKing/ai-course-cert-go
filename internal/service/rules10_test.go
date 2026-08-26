package service

import "testing"

func TestRule200(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule200(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 200: %d", r)
		}
	}
}
func TestRule201(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule201(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 201: %d", r)
		}
	}
}
func TestRule202(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule202(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 202: %d", r)
		}
	}
}
func TestRule203(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule203(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 203: %d", r)
		}
	}
}
func TestRule204(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule204(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 204: %d", r)
		}
	}
}
func TestRule205(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule205(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 205: %d", r)
		}
	}
}
func TestRule206(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule206(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 206: %d", r)
		}
	}
}
func TestRule207(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule207(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 207: %d", r)
		}
	}
}
func TestRule208(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule208(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 208: %d", r)
		}
	}
}
func TestRule209(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule209(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 209: %d", r)
		}
	}
}
