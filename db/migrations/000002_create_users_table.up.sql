CREATE TABLE "users" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "fullname" varchar(255),
    "email" varchar(255) UNIQUE NOT NULL,
    "password" varchar(255) NOT NULL,
    "photo_profile" varchar(255),
    "bio" text,
    "location" varchar(255),
    "role" user_role NOT NULL DEFAULT 'attendee',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz,
    "job_position" varchar(100)
);