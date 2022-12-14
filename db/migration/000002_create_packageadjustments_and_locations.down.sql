ALTER TABLE IF EXISTS packages DROP CONSTRAINT IF EXISTS "facility_location_id_fkey";

DROP TABLE IF EXISTS facility_locations CASCADE;
DROP TABLE IF EXISTS facilities CASCADE;
DROP TABLE IF EXISTS package_adj_entries CASCADE;