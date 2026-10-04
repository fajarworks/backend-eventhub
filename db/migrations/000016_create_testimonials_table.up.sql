CREATE TABLE "testimonials" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" int NOT NULL,
    "message" text,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "testimonials" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;