-- +goose Up
ALTER TABLE books ADD COLUMN author_surname varchar(64);

-- +goose Down
ALTER TABLE books DROP COLUMN author_surname;
