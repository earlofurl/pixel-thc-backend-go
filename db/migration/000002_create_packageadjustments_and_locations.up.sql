CREATE TABLE "package_adj_entries"
(
    "id"         bigserial PRIMARY KEY,
    "package_id" bigint      NOT NULL,
    "amount"     bigint      NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "facilities"
(
    "id"             bigserial PRIMARY KEY,
    "created_at"     timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"     timestamptz  NOT NULL DEFAULT (now()),
    "name"           varchar(255) NOT NULL DEFAULT '',
    "license_number" varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE "facility_locations"
(
    "id"          bigserial PRIMARY KEY,
    "created_at"  timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz  NOT NULL DEFAULT (now()),
    "name"        varchar(255) NOT NULL DEFAULT '',
    "facility_id" bigint       NOT NULL
);

CREATE INDEX ON "package_adj_entries" ("package_id");

COMMENT ON COLUMN "package_adj_entries"."amount" IS 'can be negative or positive';

ALTER TABLE "packages" ADD COLUMN "facility_location_id" bigint;

ALTER TABLE "packages"
    ADD FOREIGN KEY ("facility_location_id") REFERENCES "facility_locations" ("id");

ALTER TABLE "package_adj_entries"
    ADD FOREIGN KEY ("package_id") REFERENCES "packages" ("id");

ALTER TABLE "facility_locations"
    ADD FOREIGN KEY ("facility_id") REFERENCES "facilities" ("id");
