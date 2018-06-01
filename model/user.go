package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const PASSWORD_SALT_KEY string = "qwert"

type User struct {
	Base
	Id           uint   `gorm:"column:id; primary_key"`
	LoginName    string `gorm:"column:login_name"`
	PasswordHash string `gorm:"column:password_hash"`
}

func (user *User) String() string {
	return fmt.Sprintf("Id: %d, LoginName: %s, CreatedAt: %v", user.Id, user.LoginName, user.CreatedAt.Format(time.RFC3339))
}

func CreateUser(loginName string, password string) (user *User, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+PASSWORD_SALT_KEY), bcrypt.DefaultCost)
	result := DB.Create(&User{LoginName: loginName, PasswordHash: string(hashedPassword)})
	err = result.Error
	if err != nil {
		return nil, err
	}
	user = result.Value.(*User)
	log.Println("Create User", user)
	return
}

func UserLogin(loginName string, password string) (user *User, err error) {
	user = new(User)
	err = DB.Where(&User{LoginName: loginName}).First(user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password+PASSWORD_SALT_KEY))
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return
}
