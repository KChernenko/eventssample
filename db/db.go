package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable()
	createEventsTable()
	createRegistrationTable()
}

func createUserTable() {
	createTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create events table!")
	}

	createIndex := `CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email)`

	_, err = DB.Exec(createIndex)

	if err != nil {
		panic("Could not create index!")
	}
}

func createEventsTable() {
	createTable := `
		CREATE TABLE IF NOT EXISTS events(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create events table!")
	}

	createIndex := `CREATE INDEX IF NOT EXISTS idx_events_user_id ON events(user_id)`

	_, err = DB.Exec(createIndex)

	if err != nil {
		panic("Could not create index!")
	}
}

func createRegistrationTable() {
	createTable := `
		CREATE TABLE IF NOT EXISTS registrations(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create registrations table!")
	}

	createIndex := `CREATE INDEX IF NOT EXISTS idx_event_id_user_id ON registrations(event_id, user_id)`

	_, err = DB.Exec(createIndex)

	if err != nil {
		panic("Could not create index!")
	}
}
