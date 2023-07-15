-- +goose Up
ALTER TABLE journals DROP COLUMN IF EXISTS publisher;

-- +goose Down
ALTER TABLE journals ADD COLUMN publisher varchar(64);