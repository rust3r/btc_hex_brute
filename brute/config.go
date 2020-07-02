package brute

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Host    string `yaml:"host"`
	HEX     string `yaml:"hex"`
	Threads int    `yaml:"threads"`
	Output  string `yaml:"output"`
}

// NewConfig ...
func NewConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("File config.yml doesn't exist: %s\n", err)
	}

	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Error unmarshal yaml file: %v\n", err)
	}
	return &c
}
