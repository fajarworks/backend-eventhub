CREATE TABLE "speakers" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" varchar(255),
    "job_position" varchar(255),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);