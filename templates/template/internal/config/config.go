package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type ApplicationConfig struct {
	Config *Config
	DB     *gorm.DB
}

type Config struct {
	Server struct {
		Port string
		Host string
	}
	Database struct {
		URL  string
		Name string
	}
}

func LoadEnvironmentConfig() *Config {
	env := Config{}
	viper.SetConfigFile("appSettings.json")
	viper.SetConfigType("json")
	viper.SetConfigName("appSettings")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't find the file app config file : %v", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
