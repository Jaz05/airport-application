package loader

import (
	"testing"
)

type TestConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Database string `yaml:"database"`
}

func TestLoadConfigShouldLoadFile(t *testing.T) {
	var config *TestConfig
	LoadConfig("../../config/database.yml", &config)
	if config == nil {
		t.Fail()
	}
}

func TestLoadConfigShouldLoadProperties(t *testing.T) {
	var config *TestConfig
	LoadConfig("../../config/database.yml", &config)
	if len(config.Database) == 0 {
		t.Fail()
	}
}
