CREATE TABLE "user_community" (
    "user_id" int NOT NULL,
    "community_id" int NOT NULL,
    PRIMARY KEY ("user_id", "community_id")
);

ALTER TABLE "user_community" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "user_community" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id") DEFERRABLE INITIALLY IMMEDIATE;