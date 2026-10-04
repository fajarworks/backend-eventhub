CREATE TABLE "notification" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "type" notification_type,
    "title" varchar(255),
    "message" text,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "read_at" timestamptz,
    "user_id" int NOT NULL
);

ALTER TABLE "notification" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;