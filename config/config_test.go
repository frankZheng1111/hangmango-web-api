package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFilePath(t *testing.T) {
	var config JSONConfig
	path := config.ConfigFilePath("test")
	expectPathRune := []rune(path)
	expectPathRune = expectPathRune[len(expectPathRune)-16:]
	assert.Equal(t, "config/test.json", string(expectPathRune))
}

func TestInitConfig(t *testing.T) {
	var testConfig JSONConfig
	InitConfig(&testConfig)
	assert.Equal(t, 8080, testConfig.Server.Port)
	assert.Equal(t, "/tmp/hangmango-web-api-test.db", testConfig.GORM.Open)
	assert.Equal(t, "sqlite3", testConfig.GORM.Driver)
	assert.Equal(t, 100, testConfig.GORM.MaxOpen)
	assert.Equal(t, 10, testConfig.GORM.MaxIdle)
	assert.Equal(t, "test", testConfig.ENV)
	assert.Equal(t, "abandon", testConfig.Hangman.Dictionary[0])
}
