CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "public_key" varchar NOT NULL,
    "private_key" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
