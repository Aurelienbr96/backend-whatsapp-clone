-- Modify "contacts" table
ALTER TABLE "contacts" DROP CONSTRAINT "contacts_users_contacts", DROP COLUMN "contact_id", DROP COLUMN "user_contacts", ADD COLUMN "contact_user_id" uuid NOT NULL, ADD COLUMN "owner_id" uuid NOT NULL, ADD CONSTRAINT "contacts_users_contacts" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "contacts_users_contact_user" FOREIGN KEY ("contact_user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "contact_owner_id_contact_user_id" to table: "contacts"
CREATE UNIQUE INDEX "contact_owner_id_contact_user_id" ON "contacts" ("owner_id", "contact_user_id");
