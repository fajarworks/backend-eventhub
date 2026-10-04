CREATE TABLE "events" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "title" varchar(255) NOT NULL,
    "description" text,
    "image" varchar(255),
    "location" varchar(255),
    "capacity" int,
    "start_time" timestamptz,
    "end_time" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz,
    "community_id" int,
    "organizer_id" int NOT NULL
);

ALTER TABLE "events" ADD FOREIGN KEY ("organizer_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "events" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id") DEFERRABLE INITIALLY IMMEDIATE;