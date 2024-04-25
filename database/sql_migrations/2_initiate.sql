-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE kasir(
    id BIGINT NOT NULL PRIMARY KEY,
    nama VARCHAR(256),
    email VARCHAR(256) UNIQUE,
    password VARCHAR (256)
)

-- +migrate StatementEnd