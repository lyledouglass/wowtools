package utilities

import (
	"github.com/spf13/viper"
	"log"
)

// LoadConfig Loads viper config file
func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading Viper config: ", err)
	}
}
