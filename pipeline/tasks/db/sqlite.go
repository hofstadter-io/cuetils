package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func handleSQLiteExec(dbname, query string, args []interface{}) (string, error) {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return "", err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}

  res, err := stmt.Exec(args...)
	if err != nil {
		return "", err
	}

  affect, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	return fmt.Sprint(affect), nil
}

func handleSQLiteQuery(dbname, query string, args []interface{}) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}

	return db.Query(query)
}
