-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "cognito_uid" character varying(50) NULL,
  "first_name" character varying(255) NULL,
  "last_name" character varying(255) NULL,
  "first_name_ja" character varying(255) NULL,
  "last_name_ja" character varying(255) NULL,
  "email" character varying(255) NOT NULL,
  "role" character varying(25) NULL,
  "is_active" boolean NULL DEFAULT false,
  "is_email_verified" boolean NULL DEFAULT false,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_cognito_uid" UNIQUE ("cognito_uid")
);
-- Create index "idx_users_cognito_uid" to table: "users"
CREATE INDEX "idx_users_cognito_uid" ON "users" ("cognito_uid");
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
