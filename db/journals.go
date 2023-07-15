package db

import (
	"context"

	"task/models"
)

func (dbc *DB) InsertJournal(ctx context.Context, jrnl models.Journal) error {
	query := `
		INSERT INTO journals(journal_title, issue_year)
		VALUES
		($1, $2)`
	rows, err := dbc.Query(ctx, query, jrnl.JournalTitle, jrnl.Issue)
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func (dbc *DB) InsertCountry(ctx context.Context, country string) error {
	query := `UPDATE journals SET publisher_country = $1 WHERE id = 1`
	rows, err := dbc.Conn.Query(ctx, query, country)
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}
