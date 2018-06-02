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
	assert.Equal(t, "Id: 1, LoginName: email, CreatedAt: "+now.Format(time.RFC3339), user.String())
}

func TestCreateUser(t *testing.T) {
	testseed.InitTestDB(DB)
	user, _ := CreateUser("test", "pass")
	assert.Equal(t, "test", user.LoginName)
}

func TestUserLogin(t *testing.T) {
	testseed.InitTestDB(DB)
	CreateUser("test", "pass")
	var err error

	_, err = UserLogin("test", "wrongPass")
	assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())

	_, err = UserLogin("wrongTest", "wrongPass")
	assert.Equal(t, "record not found", err.Error())

	user, _ := UserLogin("test", "pass")
	assert.Equal(t, "test", user.LoginName)
}
