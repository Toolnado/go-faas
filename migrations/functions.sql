CREATE TABLE functions (
    id          UUID        PRIMARY KEY,
    name        TEXT        NOT NULL,
    version     INT         NOT NULL DEFAULT 1,
    wasm        BYTEA       NOT NULL,          
    author      TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT name_version UNIQUE(name, version)
);