package config

import (
	"os"
	"strconv"
	"time"
)

func String(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}
func Int(name string, def int) int {
	v, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		return def
	}
	return v
}
func Duration(name string, def time.Duration) time.Duration {
	v, err := time.ParseDuration(os.Getenv(name))
	if err != nil {
		return def
	}
	return v
}
func (c Config) Validate() error {
	if c.Addr == "" || c.DBPath == "" {
		return os.ErrInvalid
	}
	return nil
}
