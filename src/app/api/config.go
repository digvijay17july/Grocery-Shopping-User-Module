package api

import "os"

type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	PortNo   int
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "cfwldagsmukpvm",
			Password: "f5ea03edbf45c3ec346bb784c5b4e54435b1cc38615493d48684e485d8b00d18",
			Name:     "d1e82sinabvrsv",
			Host:     os.Getenv("DATABASE_URL"),
			Charset:  "utf8",
			PortNo:   5432,
		},
	}
}
