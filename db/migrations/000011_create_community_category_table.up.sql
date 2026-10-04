CREATE TABLE "community_category" (
    "category_id" int NOT NULL,
    "community_id" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY ("community_id", "category_id")
);

ALTER TABLE "community_category" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "community_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") DEFERRABLE INITIALLY IMMEDIATE;