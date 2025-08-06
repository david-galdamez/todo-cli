package database

import (
	"database/sql"
)

func MigrateDatabase(dbConn *DBConnection) (sql.Result, error) {
	result, err := dbConn.DB.Exec(`
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER NOT NULL PRIMARY KEY,
			todo TEXT NOT NULL,
			is_completed BOOL NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP,
			due_to TIMESTAMP
		);`)

	return result, err
}

func DropDatabase(dbConn *DBConnection) (sql.Result, error) {

	result, err := dbConn.DB.Exec(`
		DROP TABLE todo;`)

	return result, err
}
