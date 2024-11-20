-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "is_verified" boolean NOT NULL DEFAULT false;
