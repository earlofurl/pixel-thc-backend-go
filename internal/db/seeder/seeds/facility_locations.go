package seeds

import "github.com/rs/zerolog/log"

// facilityLocationsSeed seeds the facility_locations table with standard set of facility locations
func (s Seed) facilityLocationsSeed() {
	var err error

	_, err = s.tx.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 1", 1)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 1")
	}

	_, err = s.tx.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 2", 1)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 2")
	}

	_, err = s.tx.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 3", 1)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 3")
	}
}
