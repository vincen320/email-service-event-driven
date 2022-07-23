package app

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")    // name of config file (without extension)
	config.SetConfigType("env")       // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("./config/") // path to look for the config file i
	config.ReadInConfig()
	return config
}
