-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE detail_transaksi(
    transaksi_id BIGINT REFERENCES transaksi (id) ,
    menu_id BIGINT REFERENCES menu (id),
    quantity int,
)

-- +migrate StatementEnd