package config

type Redis struct {
	Addr     string
	Port     int
	Password string
}

var RedisConfig = new(Redis)
