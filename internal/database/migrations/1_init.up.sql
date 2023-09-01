CREATE TABLE "users"(
    "user_id" SERIAL PRIMARY KEY,
    "username" TEXT NOT NULL UNIQUE,
    "hashed_password" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "last_login" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "roles"(
    "role_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) CHECK ("name" IN('ADMIN', 'USER')) NOT NULL
);

CREATE TABLE "user_roles"(
    "user_id" INTEGER NOT NULL,
    "role_id" INTEGER NOT NULL,
    PRIMARY KEY ("user_id", "role_id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id"),
    FOREIGN KEY ("role_id") REFERENCES "roles"("role_id")
);

CREATE TABLE "tokens"(
    "token_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "type" VARCHAR(255) CHECK ("type" IN('ACCESS', 'REFRESH')) NOT NULL,
    "value" TEXT NOT NULL,
    "expiry_at" TIMESTAMPTZ NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users"("user_id")
);
