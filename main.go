package main

import (
	"context"
	"log"

	"task/db"
	"task/models"

	"github.com/caarlos0/env/v6"
)

type config struct {
	DBUser string `env:"DB_USER,required"`
	DBPass string `env:"DB_PASS,required"`
	DBHost string `env:"DB_HOST,required"`
	DBPort string `env:"DB_PORT" envDefault:"5432"`
	DBName string `env:"DB_NAME,required"`
}

func main() {
	ctx := context.Background()
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatalf("Read config error: %v\n", err)
		return
	}

	connectString := db.BuildString(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbc, err := db.New(ctx, connectString)
	if err := db.Migrate(connectString); err != nil {
		log.Fatalf("Failed to run migrations: %v\n", err)
		return
	}
	var book = *models.InputData()

	if err := dbc.InsertBook(ctx, book); err != nil {
		log.Fatalf("Failed to insert new book: %+v\n", err)
		return
	}
	surname := models.InputSurname()

	if err := dbc.InsertSurname(ctx, surname); err != nil {
		log.Fatalf("Failed to insert surname: %+v\n", err)
		return
	}

	jrnl := *models.InputJournal()

	if err := dbc.InsertJournal(ctx, jrnl); err != nil {
		log.Fatalf("Failed to insert new journal: %+v\n", err)
		return
	}

	country := models.InputCountry()

	if err := dbc.InsertCountry(ctx, country); err != nil {
		log.Fatalf("Failed to add new column: %+v\n", err)
		return
	}
}

func ReadConfig() (config, error) {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
