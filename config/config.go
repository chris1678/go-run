package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	ExtendConfig interface{}
	_cfg         *Config
)

// Config 配置集合
type Config struct {
	Application *Application `yaml:"application"`
	Ssl         *Ssl         `yaml:"ssl"`
	Logger      *Logger      `yaml:"logger"`
	Jwt         *Jwt         `yaml:"jwt"`
	Database    *Database    `yaml:"database"`
	Redis       *Redis       `yaml:"redis"`
	Extend      interface{}  `yaml:"extend"`
	Alioss      *Alioss      `yaml:"alioss"`
}

func (c Config) OnChange() {
	fmt.Println("OnChange")
}

// Init 初始化配置
func Init(filename string, cfg *Config) error {
	var err error
	viper.SetConfigFile(filename)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return err
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(cfg)
		if err != nil {
			log.Fatal(err)
		}
		cfg.OnChange()
	})
	viper.WatchConfig()
	return err
}

func Initialize(conf string) {
	_cfg = &Config{
		Application: ApplicationConfig,
		Ssl:         SslConfig,
		Logger:      LoggerConfig,
		Jwt:         JwtConfig,
		Database:    DatabaseConfig,
		Redis:       RedisConfig,
		Extend:      ExtendConfig,
		Alioss:      AliossConfig,
	}
	err := Init(conf, _cfg)
	if err != nil {
		log.Fatal(err)
	}

}
