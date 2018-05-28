package config

import (
	"encoding/json"
	"os"
)

type JSONConfig struct {
	ProjectName string `toml:"project_name"`
	Server      Server `toml:"server"`
}

type Server struct {
	Port int `toml:"port"`
}

var Config JSONConfig

func init() {
	filePath := "config/dev.json"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&Config); err != nil {
		panic(err)
	}
}
