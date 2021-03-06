package config

import (
	"bytes"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:"db"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	SSLmode  string `mapstructure:"sslmode"`
}

func Read() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		log.Print("No config file found")
	}

	viper.SetEnvPrefix("song")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}
