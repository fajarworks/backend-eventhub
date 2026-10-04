CREATE TABLE "event_category" (
    "event_id" int NOT NULL,
    "category_id" int NOT NULL,
    PRIMARY KEY ("event_id", "category_id")
);

ALTER TABLE "event_category" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "event_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") DEFERRABLE INITIALLY IMMEDIATE;