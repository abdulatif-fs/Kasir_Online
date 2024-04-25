-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE menu(
    id BIGINT NOT NULL PRIMARY KEY,
    nama VARCHAR(256),
    harga int
)

-- +migrate StatementEnd