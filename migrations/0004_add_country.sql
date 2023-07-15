-- +goose Up
ALTER TABLE journals ADD COLUMN publisher_country varchar(64);

-- +goose Down
ALTER TABLE journals DROP COLUMN publisher_country;
