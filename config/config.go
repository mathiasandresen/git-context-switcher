package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Context struct {
	Name       string `yaml:"name"`
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
	Email      string `yaml:"email"`
}

type Config struct {
	CurrentContext string    `yaml:"current_context"`
	Contexts       []Context `yaml:"contexts"`
}

func LoadConfig() (*Config, error) {
	configFile := os.Getenv("HOME") + "/.git-contexts.yaml"
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
