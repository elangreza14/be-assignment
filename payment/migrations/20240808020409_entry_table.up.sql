BEGIN
;

CREATE TABLE IF NOT EXISTS "entries" (
    "id" SERIAL PRIMARY KEY,
    "account_id" BIGINT REFERENCES accounts(id),
    "amount" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_entry_update" BEFORE
UPDATE
    ON "entries" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;