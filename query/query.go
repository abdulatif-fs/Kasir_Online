package query

import (
	"database/sql"
	// "fmt"
	"kasir_online/structs"
)

func GetMenu(db *sql.DB) (results []structs.Menu, err error) {
	sql := "SELECT * from menu"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var menu = structs.Menu{}

		err = rows.Scan(&menu.Id, &menu.Nama, &menu.Harga)
		// fmt.Println(err, menu.Id)
		if err != nil {
			panic(err)
		}

		results = append(results, menu)
	}
	return
}

func InsertMenu(db *sql.DB, menu structs.Menu) (err error) {
	sql := "INSERT INTO menu (id, nama, harga) VALUES ($1, $2, $3)"

	_, errs := db.Exec(sql, &menu.Id, &menu.Nama, &menu.Harga)

	return errs
}

func UpdateMenu(db *sql.DB, menu structs.Menu) (err error) {
	sql := "UPDATE menu SET nama=$1, harga=$2 WHERE id=$3"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err2 := stmt.Exec(menu.Nama, menu.Harga, menu.Id)

	if err2 != nil {
		panic(err2)
	}

	return
}

func DeleteMenu(db *sql.DB, menu structs.Menu) (err error) {
	sql := "DELETE FROM menu WHERE id=$1"

	_, errs := db.Exec(sql, &menu.Id)

	return errs
}

func GetKasir(db *sql.DB) (results []structs.Kasir, err error) {
	sql := "SELECT * from kasir"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var kasir = structs.Kasir{}

		err = rows.Scan(&kasir.Id, &kasir.Nama)
		if err != nil {
			panic(err)
		}

		results = append(results, kasir)
	}
	return
}

func InsertKasir(db *sql.DB, kasir structs.Kasir) (err error) {
	sql := "INSERT INTO kasir (id, nama) VALUES ($1, $2)"

	_, errs := db.Exec(sql, &kasir.Id, &kasir.Nama)

	return errs
}

func GetTransaksi(db *sql.DB) (results []structs.Transaksi, err error) {
	sql := "SELECT * from transaksi"
	// sql2 := "SELECT SUM(total) from detail_transaksi WHERE id=$1"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transaksi = structs.Transaksi{}
		// sql := "SELECT SUM(harga) from detail_transaksi WHERE transaksi_id = $1"
		// id :=

		err = rows.Scan(&transaksi.Id, &transaksi.Kasir_id, transaksi.Tanggal_transaksi, transaksi.Tanggal_transaksi)
		if err != nil {
			panic(err)
		}

		results = append(results, transaksi)
	}
	return
}
