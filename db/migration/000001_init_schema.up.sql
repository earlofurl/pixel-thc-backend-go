CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "hashed_password" varchar NOT NULL,
                         "username" varchar UNIQUE NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "first_name" varchar NOT NULL,
                         "last_name" varchar NOT NULL,
                         "phone" varchar NOT NULL,
                         "role" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "currency" varchar NOT NULL,
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
                            "username" varchar NOT NULL,
                            "refresh_token" varchar NOT NULL,
                            "user_agent" varchar NOT NULL,
                            "client_ip" varchar NOT NULL,
                            "is_blocked" boolean NOT NULL DEFAULT false,
                            "expires_at" timestamptz NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories" (
                                      "id" bigserial PRIMARY KEY,
                                      "name" varchar UNIQUE NOT NULL,
                                      "created_at" timestamptz NOT NULL DEFAULT (now()),
                                      "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "strains" (
                           "id" bigserial PRIMARY KEY,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now()),
                           "name" varchar NOT NULL,
                           "type" varchar,
                           "yield_average" decimal,
                           "terp_average_total" decimal,
                           "terp_1" varchar,
                           "terp_1_value" decimal,
                           "terp_2" varchar,
                           "terp_2_value" decimal,
                           "terp_3" varchar,
                           "terp_3_value" decimal,
                           "terp_4" varchar,
                           "terp_4_value" decimal,
                           "terp_5" varchar,
                           "terp_5_value" decimal,
                           "thc_average" decimal,
                           "total_cannabinoid_average" decimal,
                           "light_dep_2022" varchar DEFAULT false,
                           "fall_harvest_2022" varchar DEFAULT false,
                           "quantity_available" decimal
);

CREATE TABLE "retail_store_locations" (
                                          "id" bigserial PRIMARY KEY,
                                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                                          "updated_at" timestamptz NOT NULL DEFAULT (now()),
                                          "name" varchar UNIQUE NOT NULL,
                                          "address" varchar NOT NULL,
                                          "city" varchar NOT NULL,
                                          "state" varchar NOT NULL,
                                          "zip" integer NOT NULL,
                                          "latitude" float,
                                          "longitude" float,
                                          "note" text,
                                          "website" varchar,
                                          "flower" boolean DEFAULT false,
                                          "prerolls" boolean DEFAULT false,
                                          "pressed_hash" boolean DEFAULT false,
                                          "created_by" varchar
);

CREATE TABLE "items" (
                         "id" bigserial PRIMARY KEY,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now()),
                         "description" varchar,
                         "is_used" boolean DEFAULT false,
                         "item_type_id" bigint,
                         "strain_id" bigint
);

CREATE TABLE "item_types" (
                              "id" bigserial PRIMARY KEY,
                              "created_at" timestamptz NOT NULL DEFAULT (now()),
                              "updated_at" timestamptz NOT NULL DEFAULT (now()),
                              "product_form" varchar NOT NULL,
                              "product_modifier" varchar NOT NULL,
                              "uom_default" bigint,
                              "product_category_id" bigint
);

CREATE TABLE "package_tags" (
                                "id" bigserial PRIMARY KEY,
                                "created_at" timestamptz NOT NULL DEFAULT (now()),
                                "updated_at" timestamptz NOT NULL DEFAULT (now()),
                                "tag_number" varchar NOT NULL,
                                "is_assigned" boolean DEFAULT false,
                                "is_provisional" boolean DEFAULT true,
                                "is_active" boolean DEFAULT false,
                                "assigned_package_id" bigint
);

CREATE TABLE "packages" (
                            "id" bigserial PRIMARY KEY,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now()),
                            "tag_id" bigint,
                            "package_type" varchar NOT NULL,
                            "quantity" decimal,
                            "notes" text,
                            "packaged_date_time" timestamptz NOT NULL DEFAULT (now()),
                            "harvest_date_time" timestamptz,
                            "lab_testing_state" varchar,
                            "lab_testing_state_date_time" timestamptz,
                            "is_trade_sample" boolean DEFAULT false,
                            "is_testing_sample" boolean DEFAULT false,
                            "product_requires_remediation" boolean DEFAULT false,
                            "contains_remediated_product" boolean DEFAULT false,
                            "remediation_date_time" timestamptz,
                            "received_date_time" timestamptz,
                            "received_from_manifest_number" varchar,
                            "received_from_facility_license_number" varchar,
                            "received_from_facility_name" varchar,
                            "is_on_hold" boolean DEFAULT false,
                            "archived_date" timestamptz,
                            "finished_date" timestamptz,
                            "item_id" bigint,
                            "provisional_label" varchar,
                            "is_provisional" boolean DEFAULT false,
                            "is_sold" boolean DEFAULT false,
                            "ppu_default" decimal,
                            "ppu_on_order" decimal,
                            "total_package_price_on_order" decimal,
                            "ppu_sold_price" decimal,
                            "total_sold_price" decimal,
                            "packaging_supplies_consumed" boolean DEFAULT false,
                            "is_line_item" boolean DEFAULT false,
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
                             "test_name" varchar,
                             "batch_code" varchar,
                             "test_id_code" varchar,
                             "lab_facility_name" varchar,
                             "test_performed_date_time" timestamptz,
                             "overall_passed" boolean,
                             "test_type_name" varchar,
                             "test_passed" boolean,
                             "test_comment" varchar,
                             "thc_total_percent" decimal,
                             "thc_total_value" decimal,
                             "cbd_percent" decimal,
                             "cbd_value" decimal,
                             "terpene_total_percent" decimal,
                             "terpene_total_value" decimal,
                             "thc_a_percent" decimal,
                             "thc_a_value" decimal,
                             "delta9_thc_percent" decimal,
                             "delta9_thc_value" decimal,
                             "delta8_thc_percent" decimal,
                             "delta8_thc_value" decimal,
                             "thc_v_percent" decimal,
                             "thc_v_value" decimal,
                             "cbd_a_percent" decimal,
                             "cbd_a_value" decimal,
                             "cbn_percent" decimal,
                             "cbn_value" decimal,
                             "cbg_a_percent" decimal,
                             "cbg_a_value" decimal,
                             "cbg_percent" decimal,
                             "cbg_value" decimal,
                             "cbc_percent" decimal,
                             "cbc_value" decimal,
                             "total_cannabinoid_percent" decimal,
                             "total_cannabinoid_value" decimal
);

CREATE TABLE "lab_tests_packages" (
                                      "lab_test_id" bigint,
                                      "package_id" bigint
);

CREATE TABLE "uoms" (
                        "id" bigserial PRIMARY KEY,
                        "created_at" timestamptz NOT NULL DEFAULT (now()),
                        "updated_at" timestamptz NOT NULL DEFAULT (now()),
                        "name" varchar UNIQUE NOT NULL,
                        "abbreviation" varchar UNIQUE NOT NULL
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
                          "order_total" decimal,
                          "notes" varchar,
                          "status" varchar NOT NULL,
                          "customer_name" varchar NOT NULL
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
