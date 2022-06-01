package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(home)
	viper.SetConfigName("wowtools-cli")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
