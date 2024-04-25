package query

import (
	"database/sql"
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
