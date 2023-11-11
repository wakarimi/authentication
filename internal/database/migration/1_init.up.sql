CREATE TYPE AccountRole AS ENUM ('ADMIN', 'USER');

CREATE TABLE "accounts"
(
    "id"              SERIAL PRIMARY KEY,
    "username"        TEXT UNIQUE NOT NULL,
    "hashed_password" TEXT        NOT NULL,
    "created_at"      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_login"      TIMESTAMPTZ
);

CREATE TABLE "account_roles"
(
    "account_id" INTEGER     NOT NULL,
    "role"       AccountRole NOT NULL,
    PRIMARY KEY ("account_id", "role"),
    FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

CREATE TABLE "devices"
(
    "id"          SERIAL PRIMARY KEY,
    "account_id"  INTEGER NOT NULL,
    "fingerprint" TEXT,
    FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

CREATE TABLE "refresh_tokens"
(
    "id"         SERIAL PRIMARY KEY,
    "device_id"  INTEGER   NOT NULL,
    "token"      TEXT      NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "expires_at" TIMESTAMP NOT NULL,
    FOREIGN KEY ("device_id") REFERENCES "devices" ("id")
);
