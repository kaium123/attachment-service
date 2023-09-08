package main

import (
	"attachment/command"
	"attachment/common/logger"
	"attachment/config"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func readConfig() {
	var err error

	viper.SetConfigFile(".env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Override config parameters from environment variables if specified
	err = viper.Unmarshal(&config.Config)
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func main() {
	readConfig()
	raventClient := logger.NewRavenClient()
	logger.NewLogger(raventClient)
	command.Execute()
}
