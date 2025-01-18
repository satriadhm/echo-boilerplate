package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  string `yaml:"readTimeout"`
		WriteTimeout string `yaml:"writeTimeout"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
