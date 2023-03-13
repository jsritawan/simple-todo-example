CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "note" text NOT NULL,
  "completed" boolean NOT NULL DEFAULT false,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz,
  "delete_at" timestamptz
);
