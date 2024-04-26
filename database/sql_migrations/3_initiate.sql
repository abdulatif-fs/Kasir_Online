-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transaksi(
    id BIGINT NOT NULL PRIMARY KEY,
    tanggal TIMESTAMP,
    kasir_id int REFERENCES kasir (id)
)

-- +migrate StatementEnd