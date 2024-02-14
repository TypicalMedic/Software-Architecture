package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

// структура конфигурации (yaml файла)
type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Port      string `yaml:"port"`
	Host      string `yaml:"host"`
	OuterHost string `yaml:"outer_host"`
}
type Database struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"dbName"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func NewConfig(configfile string) (*Config, error) {
	f, err := os.Open(configfile)
	if err != nil {
		return &Config{}, err
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return &Config{}, err
	}
	return &config, nil
}
