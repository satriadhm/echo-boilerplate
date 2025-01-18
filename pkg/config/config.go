package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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
	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
	Middlewares struct {
		EnableRequestLogging bool `yaml:"enableRequestLogging"`
	} `yaml:"middlewares"`
}

// LoadConfig loads YAML configuration from a file.
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

// ConnectDatabase establishes a PostgreSQL database connection.
func ConnectDatabase(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	return sql.Open("postgres", dsn)
}
