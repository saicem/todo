package config

type Session struct {
	MaxAge int    `toml:"maxage"`
	Domain string `toml:"domain"`
}
