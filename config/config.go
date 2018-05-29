package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
	InitConfig(&Config)
}

func generateWorkPath() string {
	env := os.Getenv("GOENV")
	workDir := os.Getenv("GOWORKDIR")
	return filepath.Join(workDir, "config", fmt.Sprintf("%s.json", env))
}

func InitConfig(config *JSONConfig) {
	filePath := generateWorkPath()
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		panic(err)
	}
}
