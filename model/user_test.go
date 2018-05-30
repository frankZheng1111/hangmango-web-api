package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserString(t *testing.T) {
	now := time.Now()
	user := *new(User)
	user.Id = 1
	user.Email = "email"
	user.CreatedAt = now
	assert.Equal(t, user.String(), "Id: 1, Email: email, CreatedAt: "+now.Format(time.RFC3339))
}

func TestCreateUser(t *testing.T) {
	user, _ := CreateUser("test", "pass")
	assert.Equal(t, int(user.Id), 1)
	assert.Equal(t, user.Email, "test")
}
