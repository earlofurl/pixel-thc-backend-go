package main

import (
	"database/sql"
	"pixel-thc-backend-go/internal/db/seeder/seeds"

	"github.com/danvergara/seeder"
	"github.com/rs/zerolog/log"

	// postgres driver
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/pixel_thc_dev?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
	}

	// Create a new seeder
	s := seeds.NewSeed(db)

	if err := seeder.Execute(s); err != nil {
		log.Fatal().Err(err).Msg("Error seeding database")
	}
}
