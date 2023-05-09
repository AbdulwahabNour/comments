
CREATE TABLE IF NOT EXISTS users(
    "id" bigserial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE IF NOT EXISTS comments(
    "id" uuid,
    "user_id" Int NOT NULL,
    "body" text,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    FOREIGN KEY (user_id) REFERENCES users(id)
);

 