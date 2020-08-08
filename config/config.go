package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port   string `json:"port"`
	DBType string `json:"dbType"`
	DBPath string `json:"dbPath"`
}

var C Config

func LoadFile(filename string) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(file), &C)
}
