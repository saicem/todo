package config

type Redis struct {
	Addr            string `toml:"addr"`
	CaptchaDuration string `toml:"captchaduration"`
}
