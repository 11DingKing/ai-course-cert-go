package service

import "testing"

func TestRule100(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule100(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 100: %d", r)
		}
	}
}
func TestRule101(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule101(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 101: %d", r)
		}
	}
}
func TestRule102(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule102(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 102: %d", r)
		}
	}
}
func TestRule103(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule103(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 103: %d", r)
		}
	}
}
func TestRule104(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule104(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 104: %d", r)
		}
	}
}
func TestRule105(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule105(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 105: %d", r)
		}
	}
}
func TestRule106(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule106(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 106: %d", r)
		}
	}
}
func TestRule107(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule107(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 107: %d", r)
		}
	}
}
func TestRule108(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule108(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 108: %d", r)
		}
	}
}
func TestRule109(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule109(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 109: %d", r)
		}
	}
}
