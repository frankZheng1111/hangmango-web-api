package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const PROJECT_NAME string = "hangmango-web-api"

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

func ConfigFilePath() string {
	var projectGoPath string
	env := os.Getenv("GOENV")
	// 在GOPATH 寻找项目存在的那一条路径
	//
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, goPath := range goPaths {
		if _, err := os.Stat(filepath.Join(goPath, "src", PROJECT_NAME)); err == nil {
			projectGoPath = goPath
			break
		}
	}
	return filepath.Join(projectGoPath, "src", PROJECT_NAME, "config", fmt.Sprintf("%s.json", env))
}

func InitConfig(config *JSONConfig) {
	filePath := ConfigFilePath()
	fmt.Println(filePath)
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
