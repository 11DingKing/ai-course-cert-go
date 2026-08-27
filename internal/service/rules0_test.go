package service

import "testing"

func TestRule0(t *testing.T) {
	cases := []int{-10, 0, 0, 50, 150}
	for _, v := range cases {
		r := Rule0(v)
		if r < 0 || r > 100 {
			t.Fatalf("rule 0: %d", r)
		}
	}
}
func TestRule1(t *testing.T) {
	cases := []int{-10, 0, 1, 51, 151}
	for _, v := range cases {
		r := Rule1(v)
		if r < 1 || r > 101 {
			t.Fatalf("rule 1: %d", r)
		}
	}
}
func TestRule2(t *testing.T) {
	cases := []int{-10, 0, 2, 52, 152}
	for _, v := range cases {
		r := Rule2(v)
		if r < 2 || r > 102 {
			t.Fatalf("rule 2: %d", r)
		}
	}
}
func TestRule3(t *testing.T) {
	cases := []int{-10, 0, 3, 53, 153}
	for _, v := range cases {
		r := Rule3(v)
		if r < 3 || r > 103 {
			t.Fatalf("rule 3: %d", r)
		}
	}
}
func TestRule4(t *testing.T) {
	cases := []int{-10, 0, 4, 54, 154}
	for _, v := range cases {
		r := Rule4(v)
		if r < 4 || r > 104 {
			t.Fatalf("rule 4: %d", r)
		}
	}
}
func TestRule5(t *testing.T) {
	cases := []int{-10, 0, 5, 55, 155}
	for _, v := range cases {
		r := Rule5(v)
		if r < 5 || r > 105 {
			t.Fatalf("rule 5: %d", r)
		}
	}
}
func TestRule6(t *testing.T) {
	cases := []int{-10, 0, 6, 56, 156}
	for _, v := range cases {
		r := Rule6(v)
		if r < 6 || r > 106 {
			t.Fatalf("rule 6: %d", r)
		}
	}
}
func TestRule7(t *testing.T) {
	cases := []int{-10, 0, 7, 57, 157}
	for _, v := range cases {
		r := Rule7(v)
		if r < 7 || r > 107 {
			t.Fatalf("rule 7: %d", r)
		}
	}
}
func TestRule8(t *testing.T) {
	cases := []int{-10, 0, 8, 58, 158}
	for _, v := range cases {
		r := Rule8(v)
		if r < 8 || r > 108 {
			t.Fatalf("rule 8: %d", r)
		}
	}
}
func TestRule9(t *testing.T) {
	cases := []int{-10, 0, 9, 59, 159}
	for _, v := range cases {
		r := Rule9(v)
		if r < 9 || r > 109 {
			t.Fatalf("rule 9: %d", r)
		}
	}
}
