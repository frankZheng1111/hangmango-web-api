package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20180528154719(txn *sql.Tx) {
	if _, err := txn.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT,
		win_count INT NOT NULL DEFAULT 0,
		finish_count INT NOT NULL DEFAULT 0,
		win_rate FLOAT,
		version INT NOT NULL DEFAULT 0,
		login_name VARCHAR(50) NOT NULL UNIQUE,
		password_hash VARCHAR(100) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	CREATE INDEX login_name_index
	ON users (login_name)
	`); err != nil {
		panic(err)
	}
	if _, err := txn.Exec(`
	ALTER TABLE users CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
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
