CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "title" text,
  "completed" boolean,
  "create_at" timestamptz DEFAULT (now()),
  "update_at" timestamptz,
  "delete_at" timestamptz
);
