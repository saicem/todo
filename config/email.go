package config

type Email struct {
	ServerHost   string `toml:"serverhost"`
	ServerPort   int    `toml:"serverport"`
	FromEmail    string `toml:"fromemail"`
	FromPassword string `toml:"frompassword"`
}
