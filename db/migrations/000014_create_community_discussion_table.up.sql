CREATE TABLE "community_discussion" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" int NOT NULL,
    "community_id" int NOT NULL,
    "message" text,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "community_discussion" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "community_discussion" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id") DEFERRABLE INITIALLY IMMEDIATE;