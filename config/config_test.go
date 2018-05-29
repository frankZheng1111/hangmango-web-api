package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateConfigPath(t *testing.T) {
	expectPathRune := []rune(generateConfigPath())
	expectPathRune = expectPathRune[len(expectPathRune)-16:]
	assert.Equal(t, string(expectPathRune), "config/test.json")
}
