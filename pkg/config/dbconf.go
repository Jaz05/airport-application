package config

import (
	"airport/pkg/loader"
	"fmt"
	"log"
	"os"
)

type DBConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Database string `yaml:"database"`
}

var config *DBConfig

func GetDBConfig() *DBConfig {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	if config == nil {
		loader.LoadConfig("config/database.yml", &config)
	}
	return config
}
