package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type AppConfig struct {
	Env *Env
	DB  *gorm.DB
}

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	DBName     string `mapstructure:"DB_NAME"`
	JwtKey     string `mapstructure:"JWT_KEY"`
	RedisUrl   string `mapstructure:"REDIS_URL"`
}

func LoadEnvironmentConfig() *Env {
	env := Env{}
	viper.SetConfigFile("app_config.json")
	viper.SetConfigType("json")
	viper.SetConfigName("app_config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't find the file .env : %v", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
