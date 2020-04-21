package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
}

func NewConfiguration() *Configurations {
	//viper is used to read a config.yml file where is specified the server's port and the database configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	//Create a variable for the interface Configurations, later it will be used to unmarshal the yml
	var config Configurations
	//Read the config file. Check for errors also.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	//decode the yml into the interface
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return &config
}
