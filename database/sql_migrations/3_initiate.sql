-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transaksi(
    id BIGINT NOT NULL PRIMARY KEY,
    tanggal TIMESTAMP,
    id_kasir int REFERENCES kasir (id)
)

-- +migrate StatementEnd