package seeds

import "github.com/rs/zerolog/log"

// FacilitiesSeed seeds the facilities table with standard set of facilities
func (s Seed) FacilitiesSeed() {
	var err error

	_, err = s.db.Exec(`INSERT INTO facilities (name, license_number) VALUES ($1, $2)`, "Facility 1", "123456")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facilities table at Facility 1")
	}

	_, err = s.db.Exec(`INSERT INTO facilities (name, license_number) VALUES ($1, $2)`, "Facility 2", "654321")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facilities table at Facility 2")
	}

	_, err = s.db.Exec(`INSERT INTO facilities (name, license_number) VALUES ($1, $2)`, "Facility 3", "987654")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding facilities table at Facility 3")
	}
}
