package config

import "os"

type Config struct{ Addr, DBPath, SessionTTL string }

func Load() Config {
	c := Config{Addr: ":8080", DBPath: "course.db", SessionTTL: "24h"}
	if v := os.Getenv("ADDR"); v != "" {
		c.Addr = v
	}
	if v := os.Getenv("DB_PATH"); v != "" {
		c.DBPath = v
	}
	if v := os.Getenv("SESSION_TTL"); v != "" {
		c.SessionTTL = v
	}
	return c
}
