package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port string `json:"port"`
}

var config Config

func LoadFile(filename string) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(file), &config)

	fmt.Println(config.Port)
}

func GetHTTPPort() string {
	return config.Port
}
