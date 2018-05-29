package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFilePath(t *testing.T) {
	expectPathRune := []rune(ConfigFilePath())
	expectPathRune = expectPathRune[len(expectPathRune)-16:]
	assert.Equal(t, string(expectPathRune), "config/test.json")
}

func TestInitConfig(t *testing.T) {
	var testConfig JSONConfig
	InitConfig(&testConfig)
	assert.Equal(t, testConfig.Server.Port, 8080)
}
