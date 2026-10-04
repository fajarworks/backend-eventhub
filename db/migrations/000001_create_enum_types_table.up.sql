CREATE TYPE "user_role" AS ENUM (
    'attendee',
    'organizer',
    'admin'
);

CREATE TYPE "event_format" AS ENUM (
    'online',
    'in_person'
);

CREATE TYPE "notification_type" AS ENUM (
    'events',
    'community'
);

SELECT typname FROM pg_type WHERE typname IN ('user_role', 'event_format', 'notification_type');