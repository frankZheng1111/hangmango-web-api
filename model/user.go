package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hangmango-web-api/lib"
	"log"
	"time"
)

const PASSWORD_SALT_KEY string = "qwert"

type User struct {
	Base
	Id           int64     `gorm:"column:id; primary_key"`
	LoginName    string    `gorm:"column:login_name"`
	PasswordHash string    `gorm:"column:password_hash"`
	WinCount     int32     `gorm:"column:win_count"`
	FinishCount  int32     `gorm:"column:finish_count"`
	WinRate      *float32  `gorm:"column:win_rate"`
	Version      int       `gorm:"column:version"`
	Hangmen      []Hangman `gorm:"ForeignKey:UserId;AssociationForeignKey:Id"`
}

var UserSnowflake *lib.Snowflake

func init() {
	UserSnowflake = lib.NewSnowflake()
}

func (user *User) String() string {
	return fmt.Sprintf("Id: %d, LoginName: %s, CreatedAt: %v", user.Id, user.LoginName, user.CreatedAt.Format(time.RFC3339))
}

func (user *User) UpdateScore(isWin bool) {
	var rowsAffected int64 = 0
	for rowsAffected == 0 {
		currentVersion := user.Version
		newVersion := currentVersion + 1
		newFinishCount := user.FinishCount + 1
		newWinCount := user.WinCount
		if isWin {
			newWinCount++
		}
		var winRate float32 = float32(newWinCount) / float32(newFinishCount)
		rowsAffected = DB.Model(user).
			Where("version = ?", currentVersion).
			Updates(User{FinishCount: newFinishCount, WinCount: newWinCount, Version: newVersion, WinRate: &winRate}).
			RowsAffected
		if rowsAffected == 0 {
			DB.Where(&User{Id: user.Id}).First(user)
		}
	}
	return
}

func (user *User) HangmenById(id int64) (hangman *Hangman, err error) {
	hangman = new(Hangman)
	err = DB.Where(&Hangman{Id: id, UserId: user.Id}).First(&hangman).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserById(id int64) (user *User, err error) {
	user = new(User)
	// if not find record "First"  will return error and "Find" will not
	err = DB.Where(&User{Id: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

func CreateUser(loginName string, password string) (user *User, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+PASSWORD_SALT_KEY), bcrypt.DefaultCost)
	result := DB.Create(&User{
		Id:           UserSnowflake.Id(),
		LoginName:    loginName,
		PasswordHash: string(hashedPassword),
	})
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
	return
}

func GetBestUsers(paginate *Paginate) (count int64, users []*User) {
	users = []*User{}
	limit, offset := paginate.ParseToLimitAndOffset()
	err := DB.
		Order("win_rate desc").
		Offset(offset).Limit(limit).Find(&users).
		Offset(-1).Limit(-1).Count(&count).
		Error
	if err != nil {
		panic(err)
	}

	return
}
