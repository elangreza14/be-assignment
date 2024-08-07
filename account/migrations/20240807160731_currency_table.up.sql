BEGIN
;

CREATE TABLE IF NOT EXISTS "currencies" (
    "code" VARCHAR PRIMARY KEY,
    "description" VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_token_update" BEFORE
UPDATE
    ON "currencies" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;