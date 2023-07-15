package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const mgrPath = "./migrations"

type DB struct {
	*pgx.Conn
}

func BuildString(user, pass, host, port, name string) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=10",
		user,
		pass,
		host,
		port,
		name,
	)
}

func New(ctx context.Context, ConnectString string) (*DB, error) {
	cnct, err := pgx.Connect(ctx, ConnectString)
	if err != nil {
		return nil, err
	}

	if err := cnct.Ping(ctx); err != nil {
		return nil, errors.New("Failed to ping DataBase")
	}

	return &DB{cnct}, nil
}
