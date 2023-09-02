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
