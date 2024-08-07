BEGIN
;

CREATE SEQUENCE IF NOT EXISTS account_seq INCREMENT 1 MINVALUE 10000;

CREATE TABLE IF NOT EXISTS "accounts" (
    "id" BIGINT PRIMARY KEY DEFAULT nextval('account_seq'),
    "user_id" uuid REFERENCES users(id),
    "currency_code" VARCHAR REFERENCES currencies(code),
    "product_id" BIGINT REFERENCES products(id),
    "balance" BIGINT NOT NULL,
    "name" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_token_update" BEFORE
UPDATE
    ON "accounts" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;