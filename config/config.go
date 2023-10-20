package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Configuration struct to match the YAML structure
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DataBase string `yaml:"database"`
	}
	App struct {
		Name   string `yaml:"name"`
		Port   int    `yaml:"port"`
		JWTKey string `yaml:"jwtKey"`
	}
	Log struct {
		Path string `yaml:"path"`
	}
}

var config Config

func Load(file string) {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
		panic(err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
		panic(err)
	}
}

func GetConfig() Config {
	return config
}
