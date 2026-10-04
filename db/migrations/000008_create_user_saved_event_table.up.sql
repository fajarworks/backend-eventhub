CREATE TABLE "user_saved_event" (
    "user_id" int NOT NULL,
    "event_id" int NOT NULL,
    PRIMARY KEY ("user_id", "event_id")
);

ALTER TABLE "user_saved_event" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "user_saved_event" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") DEFERRABLE INITIALLY IMMEDIATE;