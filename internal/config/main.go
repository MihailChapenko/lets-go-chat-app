package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var config Config

//Config describe common config struct
type Config struct {
	Server *Server `json:"server" yaml:"server"`
	DB     *DB     `json:"db" yaml:"db"`
}

//Server describe server config struct
type Server struct {
	Port string `json:"port" yaml:"port"`
}

type DB struct {
	Dialect    string `json:"dialect" yaml:"dialect"`
	DataSource string `json:"datasource" yaml:"datasource"`
	Directory  string `json:"dir" yaml:"dir"`
	Table      string `json:"table" yaml:"table"`
}

//Init read and parse config from yaml
func Init(path string) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

//Get returns config instance
func Get() Config {
	return config
}
