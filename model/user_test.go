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

func TestUpdateScore(t *testing.T) {
	InitTestDB()
	user := new(User)
	DB.Where(User{Id: 1}).First(user)
	user.UpdateScore(true)
	assert.Equal(t, float32(1), user.WinRate)
	user.UpdateScore(false)
	assert.Equal(t, float32(1)/float32(2), user.WinRate)
}

func TestCreateUser(t *testing.T) {
	InitTestDB()
	user, _ := CreateUser("test", "pass")
	assert.Equal(t, "test", user.LoginName)
}

func TestUserLogin(t *testing.T) {
	InitTestDB()
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

func TestGetUserById(t *testing.T) {
	InitTestDB()
	user, _ := GetUserById(1)
	assert.Equal(t, "test-user-name", user.LoginName)
	assert.Equal(t, int64(1), user.Id)

	_, err := GetUserById(2)
	assert.NotNil(t, err)
}

func TestHangById(t *testing.T) {
	InitTestDB()
	user, _ := GetUserById(1)
	hangman, _ := user.HangmenById(1)
	assert.Equal(t, int64(1), hangman.Id)
}
