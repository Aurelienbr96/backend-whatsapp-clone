-- Create "chats" table
CREATE TABLE "chats" ("id" uuid NOT NULL, "user_chats" uuid NULL, PRIMARY KEY ("id"));
-- Create "messages" table
CREATE TABLE "messages" ("id" uuid NOT NULL, "content" character varying NOT NULL, "created_at" timestamptz NOT NULL, "read" boolean NOT NULL DEFAULT false, "chat_messages" uuid NOT NULL, "user_messages" uuid NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "phone_number" character varying NOT NULL, "chat_users" uuid NULL, PRIMARY KEY ("id"));
-- Create index "users_phone_number_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_number_key" ON "users" ("phone_number");
-- Modify "chats" table
ALTER TABLE "chats" ADD CONSTRAINT "chats_users_chats" FOREIGN KEY ("user_chats") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "messages" table
ALTER TABLE "messages" ADD CONSTRAINT "messages_chats_messages" FOREIGN KEY ("chat_messages") REFERENCES "chats" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "messages_users_messages" FOREIGN KEY ("user_messages") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "users" table
ALTER TABLE "users" ADD CONSTRAINT "users_chats_users" FOREIGN KEY ("chat_users") REFERENCES "chats" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
