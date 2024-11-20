-- Modify "users" table
ALTER TABLE "users" DROP COLUMN "email", DROP COLUMN "password";
-- Rename a column from "name" to "username"
ALTER TABLE "users" RENAME COLUMN "name" TO "username";
