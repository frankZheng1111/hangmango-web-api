package model

import (
	"github.com/stretchr/testify/assert"
	"hangmango-web-api/testseed"
	"testing"
	"time"
)

func TestUserString(t *testing.T) {
	now := time.Now()
	user := *new(User)
	user.Id = 1
	user.LoginName = "email"
	user.CreatedAt = now
	assert.Equal(t, user.String(), "Id: 1, LoginName: email, CreatedAt: "+now.Format(time.RFC3339))
}

func TestCreateUser(t *testing.T) {
	testseed.InitTestDB(DB)
	user, _ := CreateUser("test", "pass")
	assert.Equal(t, user.LoginName, "test")
}
