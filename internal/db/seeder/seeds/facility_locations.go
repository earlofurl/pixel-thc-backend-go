package seeds

import "github.com/rs/zerolog/log"

// FacilityLocationsSeed seeds the facility_locations table with standard set of facility locations
func (s Seed) FacilityLocationsSeed() {
	var err error

	_, err = s.db.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 1", "1")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 1")
	}

	_, err = s.db.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 2", "6")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 2")
	}

	_, err = s.db.Exec(`INSERT INTO facility_locations (name, facility_id) VALUES ($1, $2)`, "Location 3", "9")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facility locations table at Location 3")
	}
}
