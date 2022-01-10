package authHandler

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration Struct
type Config struct {
	Users []struct {
		ID       int    `yaml:"Id"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Services []struct {
			Name string `yaml:"name"`
			Port int    `yaml:"port"`
		} `yaml:"services"`
	} `yaml:"Users"`
}

const (
	ConfigFile = "config.yml"
)

var Configuration Config

func Load_config() {
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &Configuration)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
	fmt.Println(Configuration)

}
