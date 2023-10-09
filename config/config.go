package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Address string
}

var C Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/bililink-converter/")
	viper.AddConfigPath("$HOME/.bililink-converter")
	viper.AddConfigPath("$HOME/.config/bililink-converter")
	viper.SetConfigType("toml")
	viper.SetDefault("address", "127.0.0.1:39080")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	C.Address = viper.GetString("address")
}
