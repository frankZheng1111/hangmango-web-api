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
	user.LoginName = "email"
	user.CreatedAt = now
	assert.Equal(t, "Id: 1, LoginName: email, CreatedAt: "+now.Format(time.RFC3339), user.String())
}

func TestCreateUser(t *testing.T) {
	InitTestDB(DB)
	user, _ := CreateUser("test", "pass")
	assert.Equal(t, "test", user.LoginName)
}

func TestUserLogin(t *testing.T) {
	InitTestDB(DB)
	CreateUser("test", "pass")
	var err error

	_, err = UserLogin("test", "wrongPass")
	assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())

	_, err = UserLogin("wrongTest", "wrongPass")
	assert.Equal(t, "record not found", err.Error())

	user, err := UserLogin("test", "pass")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "test", user.LoginName)
}
