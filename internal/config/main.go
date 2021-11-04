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
}

//Server describe server config struct
type Server struct {
	Port string `json:"port" yaml:"port"`
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