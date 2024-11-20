-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "name" DROP NOT NULL, ADD COLUMN "email" character varying NOT NULL, ADD COLUMN "password" character varying NOT NULL, ADD COLUMN "avatar" character varying NULL;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
