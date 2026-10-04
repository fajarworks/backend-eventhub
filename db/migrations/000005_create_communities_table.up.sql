CREATE TABLE "communities" (
    "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" varchar(255),
    "image" varchar(255),
    "description" text
);