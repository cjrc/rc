package main

import (
	"fmt"
	"os"

	"github.com/cjrc/race/model"
	"github.com/spf13/viper"
)

// Config represents the global configuration
type Config struct {
	DB string `mapstructure:"DB"`

	HTMLPath     string
	TemplatePath string

	Events []model.Event // The events in this regatta
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in pwd directory with name "race.yml"
		viper.AddConfigPath(pwd)
		viper.SetConfigName("race")
	}

	// bind default values to config params
	for k, v := range ConfigDefaults {
		viper.SetDefault(k, v)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Printf("Fatal error config file: %s\n", err)
		os.Exit(1)
	}

	// The command line flag overrides the env variable or config file
	if dbString != "" {
		C.DB = dbString
	}
}
