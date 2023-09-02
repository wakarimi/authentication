CREATE TYPE UserRole AS ENUM ('ADMIN', 'USER');

CREATE TABLE "users"(
    "user_id" SERIAL PRIMARY KEY,
    "username" TEXT NOT NULL UNIQUE,
    "hashed_password" TEXT NOT NULL,
    "role" UserRole NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_login" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "tokens"(
    "token_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "type" VARCHAR(255) CHECK ("type" IN('ACCESS', 'REFRESH')) NOT NULL,
    "value" TEXT NOT NULL,
    "expiry_at" TIMESTAMPTZ NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id")
);
