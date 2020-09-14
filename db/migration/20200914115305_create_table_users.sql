-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS "public"."users" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" serial NOT NULL PRIMARY KEY, 

    "name" text NOT NULL, 
    "phone" text NOT NULL,
    "password" text NOT NULL,
    "role" text NOT NULL,

    "timestamp" INTEGER
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS "public"."users" CASCADE;
