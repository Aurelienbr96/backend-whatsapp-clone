-- Create "contacts" table
CREATE TABLE "contacts" ("id" uuid NOT NULL, "contact_id" character varying NOT NULL, "user_contacts" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "contacts_users_contacts" FOREIGN KEY ("user_contacts") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
