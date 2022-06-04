package config

type Application struct {
	ReadTimeout   int
	WriterTimeout int
	VerifyTimeout int
	MaxHeader     int
	Port          int64
	Mode          string
	PrivateCert   string
	PublicCert    string
}

var ApplicationConfig = new(Application)
