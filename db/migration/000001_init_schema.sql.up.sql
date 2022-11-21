CREATE TABLE "tbl_menu_items" (
    "id" BIGSERIAL PRIMARY KEY,
    "item_cd" INT NOT NULL,
    "item_name" VARCHAR NOT NULL,
    "abv" VARCHAR NOT NULL, 
    "sort" BIGINT NOT NULL,
    "outlet_id" INT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT(''), 
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_menu_item_details" (
    "id" BIGSERIAL PRIMARY KEY,
    "item_cd_detail" INT NOT NULL,
    "item_id" INT NOT NULL,
    "group_id" INT NOT NULL,
    "category_id" INT NOT NULL,
    "size_id" INT NOT NULL,
    "cost" DECIMAL NOT NULL DEFAULT(0),
    "price" DECIMAL NOT NULL DEFAULT(0),
    "vat_id" INT NOT NULL,
    "vat" DECIMAL NOT NULL DEFAULT(0),
    "terminal_id" INT NOT NULL,
    "outlet_id" INT NOT NULL,
    "printer_id" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_vats" (
    "id" BIGSERIAL PRIMARY KEY,
    "vat_cd" INT NOT NULL,
    "vat_key" VARCHAR NOT NULL,
    "vat_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT(''), 
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_terminals" (
     "id" BIGSERIAL PRIMARY KEY,
     "terminal_cd" INT NOT NULL,
     "terminal_name" VARCHAR NOT NULL,
     "ip_address" VARCHAR NOT NULL,
     "sort" BIGINT NOT NULL,
     "outlet_cd" INT NOT NULL,
     "description" varchar NOT NULL DEFAULT (''),
     "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
     "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
     "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_outlets" (
     "id" BIGSERIAL PRIMARY KEY,
     "outlet_cd" INT NOT NULL,
     "outlet_name" VARCHAR NOT NULL,
     "ip_address" VARCHAR NOT NULL,
     "sort" BIGINT NOT NULL,
     "description" VARCHAR NOT NULL DEFAULT(''),
     "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
     "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
     "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_menu_categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "category_cd" INT NOT NULL,
    "category_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT(''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_menu_groups" (
    "id" BIGSERIAL PRIMARY KEY,
    "group_cd" INT NOT NULL,
    "group_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT(''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_menu_sizes" (
    "id" BIGSERIAL PRIMARY KEY,
    "size_cd" INT NOT NULL,
    "size_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT (''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_printers" (
    "id" BIGSERIAL PRIMARY KEY,
    "print_cd" INT NOT NULL,
    "print_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "ip_address" VARCHAR NOT NULL DEFAULT(''),
    "description" VARCHAR NOT NULL DEFAULT(''),
    "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "tbl_menu_modifies" (
    "id" BIGSERIAL PRIMARY KEY,
    "modify_cd" INT NOT NULL,
    "modify_name" VARCHAR NOT NULL,
    "sort" BIGINT NOT NULL,
    "description" VARCHAR NOT NULL DEFAULT (''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE "tbl_menu_item_modifies"(
    "id" BIGSERIAL PRIMARY KEY,
    "item_cd" INT NOT NULL,
    "modify_cd" INT NOT NULL
);


ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("terminal_id") REFERENCES "tbl_terminals" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("printer_id") REFERENCES "tbl_printers" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("outlet_id") REFERENCES "tbl_outlets" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("group_id") REFERENCES "tbl_menu_groups" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("category_id") REFERENCES "tbl_menu_categories" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("size_id") REFERENCES "tbl_menu_sizes" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("vat_id") REFERENCES "tbl_vats" ("id");
ALTER TABLE "tbl_menu_item_details" ADD FOREIGN KEY ("item_id") REFERENCES "tbl_menu_items" ("id");

CREATE INDEX ON "tbl_terminals" ("terminal_cd");
CREATE INDEX ON "tbl_outlets" ("outlet_cd");
CREATE INDEX ON "tbl_menu_item_details" ("item_cd_detail");
CREATE INDEX ON "tbl_menu_items" ("item_cd");
CREATE INDEX ON "tbl_menu_categories" ("category_cd");
CREATE INDEX ON "tbl_menu_groups" ("group_cd");
CREATE INDEX ON "tbl_menu_sizes" ("size_cd");
CREATE INDEX ON "tbl_vats" ("vat_cd");
CREATE INDEX ON "tbl_printers" ("print_cd");
CREATE INDEX ON "tbl_menu_item_modifies" ("id");
CREATE INDEX ON "tbl_menu_modifies" ("modify_cd");

COMMENT ON COLUMN "tbl_menu_item_details"."cost" IS 'must positive';
COMMENT ON COLUMN "tbl_menu_item_details"."price" IS 'must be positive';