package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Migrate(connectString string) error {
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.Up(db, mgrPath); err != nil {
		return err
	}

	return nil
}
