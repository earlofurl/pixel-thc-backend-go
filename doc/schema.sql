-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-12-07T06:30:07.626Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "hashed_password" varchar(64) NOT NULL,
  "username" varchar(64) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "phone" varchar(26) NOT NULL,
  "role" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar(64) NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar(3) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar(64) NOT NULL,
  "refresh_token" varchar(255) NOT NULL,
  "user_agent" varchar(255) NOT NULL,
  "client_ip" varchar(45) NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "strains" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) NOT NULL,
  "type" varchar(255) NOT NULL,
  "yield_average" numeric(9,6),
  "terp_average_total" numeric(9,6),
  "terp_1" varchar(255),
  "terp_1_value" numeric(9,6),
  "terp_2" varchar(255),
  "terp_2_value" numeric(9,6),
  "terp_3" varchar(255),
  "terp_3_value" numeric(9,6),
  "terp_4" varchar(255),
  "terp_4_value" numeric(9,6),
  "terp_5" varchar(255),
  "terp_5_value" numeric(9,6),
  "thc_average" numeric(9,6),
  "total_cannabinoid_average" numeric(9,6),
  "light_dep_2022" varchar DEFAULT 'false',
  "fall_harvest_2022" varchar DEFAULT 'false',
  "quantity_available" numeric(9,6)
);

CREATE TABLE "retailer_locations" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) UNIQUE NOT NULL,
  "address" varchar(255) NOT NULL,
  "city" varchar(255) NOT NULL,
  "state" varchar(255) NOT NULL,
  "zip" varchar(10) NOT NULL,
  "latitude" numeric(9,6) NOT NULL DEFAULT 0,
  "longitude" numeric(9,6) NOT NULL DEFAULT 0,
  "note" varchar(1024),
  "website" varchar(2083),
  "sells_flower" boolean DEFAULT false,
  "sells_prerolls" boolean DEFAULT false,
  "sells_pressed_hash" boolean DEFAULT false,
  "created_by" varchar(255)
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "description" varchar(255) NOT NULL DEFAULT '',
  "is_used" boolean NOT NULL DEFAULT false,
  "item_type_id" bigint NOT NULL,
  "strain_id" bigint NOT NULL
);

CREATE TABLE "item_types" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "product_form" varchar(255) NOT NULL,
  "product_modifier" varchar(255) NOT NULL,
  "uom_default" bigint NOT NULL,
  "product_category_id" bigint NOT NULL
);

CREATE TABLE "package_tags" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "tag_number" varchar(24) NOT NULL,
  "is_assigned" boolean NOT NULL DEFAULT false,
  "is_provisional" boolean NOT NULL DEFAULT true,
  "is_active" boolean NOT NULL DEFAULT false,
  "assigned_package_id" bigint
);

CREATE TABLE "packages" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "tag_id" bigint,
  "package_type" varchar(255) NOT NULL,
  "quantity" numeric(9,6),
  "notes" varchar(1024),
  "packaged_date_time" timestamptz NOT NULL DEFAULT (now()),
  "harvest_date_time" timestamptz,
  "lab_testing_state" varchar(255),
  "lab_testing_state_date_time" timestamptz,
  "is_trade_sample" boolean NOT NULL DEFAULT false,
  "is_testing_sample" boolean NOT NULL DEFAULT false,
  "product_requires_remediation" boolean NOT NULL DEFAULT false,
  "contains_remediated_product" boolean NOT NULL DEFAULT false,
  "remediation_date_time" timestamptz,
  "received_date_time" timestamptz,
  "received_from_manifest_number" varchar(255),
  "received_from_facility_license_number" varchar(255),
  "received_from_facility_name" varchar(255),
  "is_on_hold" boolean NOT NULL DEFAULT false,
  "archived_date" timestamptz,
  "finished_date" timestamptz,
  "item_id" bigint,
  "provisional_label" varchar(255),
  "is_provisional" boolean NOT NULL DEFAULT false,
  "is_sold" boolean NOT NULL DEFAULT false,
  "ppu_default" numeric(19,4),
  "ppu_on_order" numeric(19,4),
  "total_package_price_on_order" numeric(19,4),
  "ppu_sold_price" numeric(19,4),
  "total_sold_price" numeric(19,4),
  "packaging_supplies_consumed" boolean NOT NULL DEFAULT false,
  "is_line_item" boolean NOT NULL DEFAULT false,
  "order_id" bigint,
  "uom_id" bigint
);

CREATE TABLE "source_packages_child_packages" (
  "source_package_id" bigint,
  "child_package_id" bigint
);

CREATE TABLE "lab_tests" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "test_name" varchar(255),
  "batch_code" varchar(255),
  "test_id_code" varchar(255),
  "lab_facility_name" varchar(255),
  "test_performed_date_time" timestamptz,
  "overall_passed" boolean,
  "test_type_name" varchar(255),
  "test_passed" boolean,
  "test_comment" varchar(255),
  "thc_total_percent" numeric(9,6),
  "thc_total_value" numeric(9,6),
  "cbd_percent" numeric(9,6),
  "cbd_value" numeric(9,6),
  "terpene_total_percent" numeric(9,6),
  "terpene_total_value" numeric(9,6),
  "thc_a_percent" numeric(9,6),
  "thc_a_value" numeric(9,6),
  "delta9_thc_percent" numeric(9,6),
  "delta9_thc_value" numeric(9,6),
  "delta8_thc_percent" numeric(9,6),
  "delta8_thc_value" numeric(9,6),
  "thc_v_percent" numeric(9,6),
  "thc_v_value" numeric(9,6),
  "cbd_a_percent" numeric(9,6),
  "cbd_a_value" numeric(9,6),
  "cbn_percent" numeric(9,6),
  "cbn_value" numeric(9,6),
  "cbg_a_percent" numeric(9,6),
  "cbg_a_value" numeric(9,6),
  "cbg_percent" numeric(9,6),
  "cbg_value" numeric(9,6),
  "cbc_percent" numeric(9,6),
  "cbc_value" numeric(9,6),
  "total_cannabinoid_percent" numeric(9,6),
  "total_cannabinoid_value" numeric(9,6)
);

CREATE TABLE "lab_tests_packages" (
  "lab_test_id" bigint,
  "package_id" bigint
);

CREATE TABLE "uoms" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(32) NOT NULL,
  "abbreviation" varchar(16) NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "scheduled_pack_date_time" timestamptz,
  "scheduled_ship_date_time" timestamptz,
  "scheduled_delivery_date_time" timestamptz,
  "actual_pack_date_time" timestamptz,
  "actual_ship_date_time" timestamptz,
  "actual_delivery_date_time" timestamptz,
  "order_total" numeric(19,4),
  "notes" varchar(1024),
  "status" varchar(255) NOT NULL DEFAULT '',
  "customer_name" varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE "package_adjustments" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "from_package_id" bigint NOT NULL,
  "to_package_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "uom_id" bigint NOT NULL
);

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

CREATE INDEX ON "package_adjustments" ("from_package_id");

CREATE INDEX ON "package_adjustments" ("to_package_id");

CREATE INDEX ON "package_adjustments" ("from_package_id", "to_package_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

COMMENT ON COLUMN "package_adjustments"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "items" ADD FOREIGN KEY ("item_type_id") REFERENCES "item_types" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("strain_id") REFERENCES "strains" ("id");

ALTER TABLE "item_types" ADD FOREIGN KEY ("uom_default") REFERENCES "uoms" ("id");

ALTER TABLE "item_types" ADD FOREIGN KEY ("product_category_id") REFERENCES "product_categories" ("id");

ALTER TABLE "package_tags" ADD FOREIGN KEY ("assigned_package_id") REFERENCES "packages" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("tag_id") REFERENCES "package_tags" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("uom_id") REFERENCES "uoms" ("id");

ALTER TABLE "source_packages_child_packages" ADD FOREIGN KEY ("source_package_id") REFERENCES "packages" ("id");

ALTER TABLE "source_packages_child_packages" ADD FOREIGN KEY ("child_package_id") REFERENCES "packages" ("id");

ALTER TABLE "lab_tests_packages" ADD FOREIGN KEY ("lab_test_id") REFERENCES "lab_tests" ("id");

ALTER TABLE "lab_tests_packages" ADD FOREIGN KEY ("package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("from_package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("to_package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("uom_id") REFERENCES "uoms" ("id");
