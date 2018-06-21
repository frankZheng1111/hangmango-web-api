package serializer

import (
	db "hangmango-web-api/model"
	"reflect"
	"time"
)

type GuessedLetter struct {
	Letter    string    `json:"letter"`
	CreatedAt time.Time `json:"created_at"`
}

func SerializeGuessedLetter(hangmanGuessedLetters []*db.HangmanGuessedLetter) []*GuessedLetter {
	guessedLetters := []*GuessedLetter{}
	for _, hangmanGuessedLetter := range hangmanGuessedLetters {
		guessedLetter := new(GuessedLetter)
		guessedLetterType := reflect.TypeOf(guessedLetter).Elem()
		for i := 0; i < guessedLetterType.NumField(); i++ {
			field := reflect.ValueOf(guessedLetter).Elem().Field(i)
			fieldValue := reflect.ValueOf(hangmanGuessedLetter).Elem().FieldByName(guessedLetterType.Field(i).Name)
			field.Set(fieldValue)
		}
		guessedLetters = append(guessedLetters, guessedLetter)
	}

	return guessedLetters
}
