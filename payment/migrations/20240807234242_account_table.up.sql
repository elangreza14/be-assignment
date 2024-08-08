BEGIN
;

CREATE TABLE IF NOT EXISTS "accounts" (
    "id" SERIAL PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "currency_code" VARCHAR NOT NULL,
    "product_id" BIGINT NOT NULL,
    "balance" BIGINT NOT NULL,
    "name" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_account_update" BEFORE
UPDATE
    ON "accounts" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;