CREATE TABLE "categories" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" varchar(255) UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);