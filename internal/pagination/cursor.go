package pagination

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func Encode(offset int) string {
	return base64.RawURLEncoding.EncodeToString([]byte(strconv.Itoa(offset)))
}
func Decode(v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	b, e := base64.RawURLEncoding.DecodeString(v)
	if e != nil {
		return 0, e
	}
	return strconv.Atoi(strings.TrimSpace(string(b)))
}
func Next(p Page, count int) string {
	if count < p.Limit {
		return ""
	}
	return Encode(p.Offset + p.Limit)
}
