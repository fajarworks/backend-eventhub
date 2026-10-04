CREATE TABLE "user_event" (
    "user_id" int NOT NULL,
    "event_id" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz,
    PRIMARY KEY ("user_id", "event_id")
);

ALTER TABLE "user_event" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "user_event" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") DEFERRABLE INITIALLY IMMEDIATE;