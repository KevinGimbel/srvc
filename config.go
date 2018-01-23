package srvc

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var version string

// Config represents the yaml config
type Config struct {
	Headers []Header               `yaml:"headers"`
	Routes  map[string]RouteConfig `yaml:"routes"`
}

// Header defines
type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// RouteConfig represents the config for a single route
type RouteConfig struct {
	Headers []Header `yaml:"headers"`
	Content string   `yaml:"content"`
	File    string   `yaml:"file"`
}

var config Config

func init() {
	f, err := ioutil.ReadFile("./srvc.yaml")

	if err != nil {
		fmt.Println("Cannot open file srvc.yaml")
	}

	err = yaml.Unmarshal(f, &config)

	if err != nil {
		fmt.Println(err)
	}
}

// GetConfig returns the config
func GetConfig() Config {
	return config
}
