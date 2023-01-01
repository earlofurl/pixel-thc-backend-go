package seeds

import "database/sql"

type Seed struct {
	db *sql.DB
}

// NewSeed returns a Seed with a pool of connections to the database
func NewSeed(db *sql.DB) Seed {
	return Seed{
		db: db,
	}
}
