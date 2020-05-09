/*
Package config contains all the configurations for the service
*/
package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/structs"
	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	v := viper.New()
	v.SetConfigType("toml")

	if strings.ToLower(env) == "dev" {
		v.AddConfigPath("config/")
	}
	_ = os.Setenv("env", env)

	v.SetConfigName(strings.ToLower(env))
	err = v.MergeInConfig()
	if err != nil {
		log.Fatal("Error on parsing configuration file. Error " + err.Error())
	}
	fmt.Println(env)
	// Set structs default tag name
	structs.DefaultTagName = "json"
	config = v
}

// GetConfig function to expose the config object
func GetConfig() *viper.Viper {
	return config
}
