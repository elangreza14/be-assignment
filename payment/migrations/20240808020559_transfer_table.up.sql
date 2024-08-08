BEGIN
;

CREATE TABLE IF NOT EXISTS "transfers" (
    "id" SERIAL PRIMARY KEY,
    "from_account_id" BIGINT REFERENCES accounts(id),
    "to_account_id" BIGINT REFERENCES accounts(id),
    "amount" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_transfer_update" BEFORE
UPDATE
    ON "transfers" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;