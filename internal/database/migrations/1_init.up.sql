CREATE TYPE AccountRole AS ENUM ('ADMIN', 'USER');

CREATE TABLE "accounts"(
    "account_id" SERIAL PRIMARY KEY,
    "username" TEXT NOT NULL UNIQUE,
    "hashed_password" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_login" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "account_roles"(
    "account_id" INTEGER NOT NULL,
    "role" AccountRole NOT NULL,
    PRIMARY KEY ("account_id", "role"),
    FOREIGN KEY ("account_id") REFERENCES "accounts"("account_id")
);

CREATE TABLE "tokens"(
    "token_id" SERIAL PRIMARY KEY,
    "account_id" INTEGER NOT NULL,
    "type" VARCHAR(255) CHECK ("type" IN('ACCESS', 'REFRESH')) NOT NULL,
    "value" TEXT NOT NULL,
    "expiry_at" TIMESTAMPTZ NOT NULL,
    FOREIGN KEY ("account_id") REFERENCES "accounts"("account_id")
);
