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
	LoginName    string `gorm:"column:email"`
	PasswordHash string `gorm:"column:password_hash"`
}

func (user *User) String() string {
	return fmt.Sprintf("Id: %d, LoginName: %s, CreatedAt: %v", user.Id, user.LoginName, user.CreatedAt.Format(time.RFC3339))
}

func CreateUser(loginName string, password string) (user *User, err error) {
	saltKey := "qwert"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+saltKey), bcrypt.DefaultCost)
	result := DB.Create(&User{LoginName: loginName, PasswordHash: string(hashedPassword)})
	err = result.Error
	if err != nil {
		return nil, err
	}
	user = result.Value.(*User)
	log.Println("Create User", user)
	return
}
