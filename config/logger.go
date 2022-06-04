package config

type Logger struct {
	Path   string
	Level  string
	Stdout string
	Cap    uint
}

var LoggerConfig = new(Logger)
