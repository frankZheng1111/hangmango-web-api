package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20180528154719(txn *sql.Tx) {
	if _, err := txn.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED AUTO_INCREMENT,
		login_name VARCHAR(50) NOT NULL UNIQUE,
		password_hash VARCHAR(100) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	CREATE INDEX password_login_index
	ON users (login_name, password_hash)
	`); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20180528154719(txn *sql.Tx) {
	if _, err := txn.Exec(`DROP TABLE users;`); err != nil {
		panic(err)
	}
}
