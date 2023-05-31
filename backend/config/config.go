package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Server struct {
		Port 			int 		`mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		User 			string 	`mapstructure:"user"`
		Password 	string 	`mapstructure:"password"`
		Host 			string 	`mapstructure:"host"`
		Port 			int 		`mapstructure:"port"`
		DBName 		string 	`mapstructure:"dbname"`
	} `mapstructure:"database"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	Cfg = cfg

	return cfg
}

func (c *Config) LoadDBURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
}

func (c *Config) ServerPort() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}
