package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Base
	Id           uint   `gorm:"column:id; primary_key"`
	Email        string `gorm:"column:email"`
	PasswordHash string `gorm:"column:password_hash"`
}

func (user *User) String() string {
	return fmt.Sprintf("Id: %d, Email: %s, CreatedAt: %v", user.Id, user.Email, user.CreatedAt.Format(time.RFC3339))
}

func CreateUser(email string, password string) (user *User, err error) {
	saltKey := "qwert"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+saltKey), bcrypt.DefaultCost)
	result := DB.Create(&User{Email: email, PasswordHash: string(hashedPassword)})
	err = result.Error
	if err != nil {
		return nil, err
	}
	user = result.Value.(*User)
	log.Println("Create User", user)
	return
}
