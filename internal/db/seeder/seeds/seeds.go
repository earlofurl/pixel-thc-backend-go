package seeds

import "database/sql"

type Seed struct {
	tx *sql.Tx
}

// PopulateDatabase calls the seed functions in specific order
func (s Seed) PopulateDatabase() {
	s.uomsSeed()
	s.facilitiesSeed()
	s.facilityLocationsSeed()
}

// NewSeed returns a Seed with a pool of connections to the database
func NewSeed(tx *sql.Tx) Seed {
	return Seed{
		tx: tx,
	}
}
