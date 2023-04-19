
CREATE TABLE IF NOT EXISTS users(
    "id" bigserial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE IF NOT EXISTS comments(
    "id" uuid,
    "slug" text,
    "author" text,
    "body" text
);