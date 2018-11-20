package main

import (
	"database/sql"
	"os"
)

type Database struct {
	db *sql.DB
}

func (d *Database) Destroy() error {
	return os.Remove("./users.db")
}

func (d *Database) Setup() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}

	d.db = db
	
	statement, err := db.Prepare(`
	  CREATE TABLE IF NOT EXISTS users (
		  id INTEGER PRIMARY KEY,
		  username varchar(128)
	  )
	`)

	if err != nil {
		return err
	}
	
	_, err = statement.Exec()
	return err
}

// AddUser adds a user with the given username.
func (d *Database) AddUser(username string) error {
	statement, err := d.db.Prepare("INSERT INTO users (username) VALUES (?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(username)
	return err
}