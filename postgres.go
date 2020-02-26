package main

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable")
}

func newStateFromDB(db *sql.DB) (*State, error) {
	rows, err := db.Query("SELECT datname FROM pg_database")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	state := &State{}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		if name != "postgres" && !strings.HasPrefix(name, "template") {
			state.Databases = append(state.Databases, Database{Name: name})
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return state, nil
}

func createDB(db *sql.DB, dbToCreate Database) error {
	log.Infof("Creating db: %s", dbToCreate.Name)
	_, err := db.Exec("CREATE DATABASE " + dbToCreate.Name)
	return err
}

func deleteDB(db *sql.DB, dbToDelete Database) error {
	log.Infof("Deleting db: %s", dbToDelete.Name)
	_, err := db.Exec("DROP DATABASE " + dbToDelete.Name)
	return err
}
