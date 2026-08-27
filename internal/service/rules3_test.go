package service

import "testing"

func TestRule60(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule60(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 60: %d", r)
		}
	}
}
func TestRule61(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule61(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 61: %d", r)
		}
	}
}
func TestRule62(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule62(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 62: %d", r)
		}
	}
}
func TestRule63(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule63(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 63: %d", r)
		}
	}
}
func TestRule64(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule64(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 64: %d", r)
		}
	}
}
func TestRule65(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule65(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 65: %d", r)
		}
	}
}
func TestRule66(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule66(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 66: %d", r)
		}
	}
}
func TestRule67(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule67(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 67: %d", r)
		}
	}
}
func TestRule68(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule68(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 68: %d", r)
		}
	}
}
func TestRule69(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule69(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 69: %d", r)
		}
	}
}
