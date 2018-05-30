package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFilePath(t *testing.T) {
	expectPathRune := []rune(ConfigFilePath("test"))
	expectPathRune = expectPathRune[len(expectPathRune)-16:]
	assert.Equal(t, string(expectPathRune), "config/test.json")
}

func TestInitConfig(t *testing.T) {
	var testConfig JSONConfig
	InitConfig(&testConfig)
	assert.Equal(t, testConfig.Server.Port, 8080)
	assert.Equal(t, testConfig.GORM.Open, "/tmp/hangmango-web-api-test.db")
	assert.Equal(t, testConfig.GORM.Driver, "sqlite3")
	assert.Equal(t, testConfig.GORM.MaxOpen, 100)
	assert.Equal(t, testConfig.GORM.MaxIdle, 10)
	assert.Equal(t, testConfig.ENV, "test")
}
