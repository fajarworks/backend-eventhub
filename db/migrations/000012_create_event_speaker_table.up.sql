CREATE TABLE "event_speaker" (
    "event_id" int NOT NULL,
    "speaker_id" int NOT NULL,
    PRIMARY KEY ("event_id", "speaker_id")
);

ALTER TABLE "event_speaker" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "event_speaker" ADD FOREIGN KEY ("speaker_id") REFERENCES "speakers" ("id") DEFERRABLE INITIALLY IMMEDIATE;