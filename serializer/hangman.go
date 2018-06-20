package serializer

import (
	db "hangmango-web-api/model"
	"reflect"
	"time"
)

type GuessingHangman struct {
	Id   int64  `json:"id"`
	Word string `json:"word"`
	Hp   int    `json:"hp"`
}

type CompletedHangman struct {
	Id        int64     `json:"id"`
	Word      string    `json:"word"`
	Hp        int       `json:"hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CompletedHangmanResource struct {
	BaseResource
	Data []*CompletedHangman `json:"data"`
}

func SerializeGuessingHangman(hangman *db.Hangman) *GuessingHangman {
	guessingHangman := new(GuessingHangman)
	guessingHangman.Id = hangman.Id
	guessingHangman.Hp = hangman.Hp
	guessingHangman.Word = hangman.GameStr()
	return guessingHangman
}

func SerializeCompletedHangman(count int64, hangmen []*db.Hangman) *CompletedHangmanResource {
	completedHangmanResource := new(CompletedHangmanResource)
	completedHangmanResource.TotalCount = count
	completedHangmanResource.Data = []*CompletedHangman{}
	for _, hangman := range hangmen {
		completedHangman := new(CompletedHangman)
		completedHangmanType := reflect.TypeOf(completedHangman).Elem()
		for i := 0; i < completedHangmanType.NumField(); i++ {
			field := reflect.ValueOf(completedHangman).Elem().Field(i)
			fieldValue := reflect.ValueOf(hangman).Elem().FieldByName(completedHangmanType.Field(i).Name)
			field.Set(fieldValue)
		}
		completedHangmanResource.Data = append(completedHangmanResource.Data, completedHangman)
	}
	return completedHangmanResource
}
