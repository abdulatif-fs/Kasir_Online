-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE kasir(
    id BIGINT NOT NULL PRIMARY KEY,
    nama VARCHAR(256)
)

-- +migrate StatementEnd