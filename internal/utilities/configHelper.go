package utilities

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig Loads viper config file
func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading Viper config: ", err)
		Log.WithError(err).Error("Error loading Viper config")
	}
}
