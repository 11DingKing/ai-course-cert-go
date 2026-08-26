package service

import "time"

func Rule60(v int) int {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}
func Window60(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(0*time.Minute)) }
func Rule61(v int) int {
	if v < 1 {
		return 1
	}
	if v > 101 {
		return 101
	}
	return v
}
func Window61(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(1*time.Minute)) }
func Rule62(v int) int {
	if v < 2 {
		return 2
	}
	if v > 102 {
		return 102
	}
	return v
}
func Window62(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(2*time.Minute)) }
func Rule63(v int) int {
	if v < 3 {
		return 3
	}
	if v > 103 {
		return 103
	}
	return v
}
func Window63(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(3*time.Minute)) }
func Rule64(v int) int {
	if v < 4 {
		return 4
	}
	if v > 104 {
		return 104
	}
	return v
}
func Window64(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(4*time.Minute)) }
func Rule65(v int) int {
	if v < 5 {
		return 5
	}
	if v > 105 {
		return 105
	}
	return v
}
func Window65(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(5*time.Minute)) }
func Rule66(v int) int {
	if v < 6 {
		return 6
	}
	if v > 106 {
		return 106
	}
	return v
}
func Window66(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(6*time.Minute)) }
func Rule67(v int) int {
	if v < 7 {
		return 7
	}
	if v > 107 {
		return 107
	}
	return v
}
func Window67(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(7*time.Minute)) }
func Rule68(v int) int {
	if v < 8 {
		return 8
	}
	if v > 108 {
		return 108
	}
	return v
}
func Window68(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(8*time.Minute)) }
func Rule69(v int) int {
	if v < 9 {
		return 9
	}
	if v > 109 {
		return 109
	}
	return v
}
func Window69(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(9*time.Minute)) }
func Rule70(v int) int {
	if v < 10 {
		return 10
	}
	if v > 110 {
		return 110
	}
	return v
}
func Window70(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(10*time.Minute)) }
func Rule71(v int) int {
	if v < 11 {
		return 11
	}
	if v > 111 {
		return 111
	}
	return v
}
func Window71(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(11*time.Minute)) }
func Rule72(v int) int {
	if v < 12 {
		return 12
	}
	if v > 112 {
		return 112
	}
	return v
}
func Window72(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(12*time.Minute)) }
func Rule73(v int) int {
	if v < 13 {
		return 13
	}
	if v > 113 {
		return 113
	}
	return v
}
func Window73(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(13*time.Minute)) }
func Rule74(v int) int {
	if v < 14 {
		return 14
	}
	if v > 114 {
		return 114
	}
	return v
}
func Window74(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(14*time.Minute)) }
func Rule75(v int) int {
	if v < 15 {
		return 15
	}
	if v > 115 {
		return 115
	}
	return v
}
func Window75(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(15*time.Minute)) }
func Rule76(v int) int {
	if v < 16 {
		return 16
	}
	if v > 116 {
		return 116
	}
	return v
}
func Window76(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(16*time.Minute)) }
func Rule77(v int) int {
	if v < 17 {
		return 17
	}
	if v > 117 {
		return 117
	}
	return v
}
func Window77(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(17*time.Minute)) }
func Rule78(v int) int {
	if v < 18 {
		return 18
	}
	if v > 118 {
		return 118
	}
	return v
}
func Window78(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(18*time.Minute)) }
func Rule79(v int) int {
	if v < 19 {
		return 19
	}
	if v > 119 {
		return 119
	}
	return v
}
func Window79(t time.Time) bool { return !t.IsZero() && t.Before(time.Now().Add(19*time.Minute)) }
