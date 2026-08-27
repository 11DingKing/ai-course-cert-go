package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func HashPassword(v string) string { h := sha256.Sum256([]byte(v)); return hex.EncodeToString(h[:]) }
func VerifyPassword(hash, password string) bool {
	return strings.EqualFold(hash, HashPassword(password))
}
func PasswordStrength(v string) int {
	n := 0
	if len(v) >= 8 {
		n++
	}
	for _, r := range v {
		if r >= '0' && r <= '9' {
			n++
			break
		}
	}
	for _, r := range v {
		if r >= 'A' && r <= 'Z' {
			n++
			break
		}
	}
	return n
}
