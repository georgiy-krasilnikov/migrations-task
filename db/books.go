package db

import (
	"context"

	"task/models"
)

func (dbc *DB) InsertBook(ctx context.Context, book models.Book) error {
	query := `
		INSERT INTO books (book_title, book_year, author_name)
		VALUES ($1, $2, $3)`
	rows, err := dbc.Query(ctx, query, book.BookTitle, book.Data, book.AuthorName)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}

func (dbс *DB) InsertSurname(ctx context.Context, surname string) error {
	query := `UPDATE books SET author_surname = $1 WHERE id = 1`
	rows, err := dbс.Conn.Query(ctx, query, surname)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}
