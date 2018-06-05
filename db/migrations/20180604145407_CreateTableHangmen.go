package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20180604145407(txn *sql.Tx) {
	if _, err := txn.Exec(`
	CREATE TABLE IF NOT EXISTS hangmen (
		id INT UNSIGNED AUTO_INCREMENT,
		user_id INT UNSIGNED NOT NULL,
		word VARCHAR(50) NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT "PLAYING",
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	CREATE INDEX user_id_index
	ON hangmen(user_id);
	`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	CREATE INDEX user_id_status_index
	ON hangmen(user_id, status);
	`); err != nil {
		panic(err)
	}

	if _, err := txn.Exec(`
	CREATE TABLE IF NOT EXISTS hangman_guessed_letters (
		id INT UNSIGNED AUTO_INCREMENT,
		hangman_id INT UNSIGNED NOT NULL,
		letter VARCHAR(1) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	CREATE INDEX hangman_id_index
	ON hangman_guessed_letters(hangman_id);
	`); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20180604145407(txn *sql.Tx) {
	if _, err := txn.Exec(`DROP TABLE hangmen;`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`DROP TABLE hangman_guessed_letters;`); err != nil {
		panic(err)
	}
}
