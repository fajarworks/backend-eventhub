CREATE TABLE "event_discussion" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" int NOT NULL,
    "event_id" int NOT NULL,
    "message" text,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "event_discussion" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "event_discussion" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") DEFERRABLE INITIALLY IMMEDIATE;