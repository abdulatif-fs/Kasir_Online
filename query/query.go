package query

import (
	"database/sql"
	"time"

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

		err = rows.Scan(&kasir.Id, &kasir.Nama, &kasir.Email, &kasir.Password)
		if err != nil {
			panic(err)
		}

		results = append(results, kasir)
	}
	return
}

func InsertKasir(db *sql.DB, kasir structs.Kasir) (err error) {
	sql := "INSERT INTO kasir (id, nama, email, password) VALUES ($1, $2, $3, $4)"

	_, errs := db.Exec(sql, &kasir.Id, &kasir.Nama, &kasir.Email, &kasir.Password)

	return errs
}

func GetTransaksi(db *sql.DB) (results []structs.Transaksi, err error) {
	sql := "SELECT * from transaksi"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transaksi = structs.Transaksi{}

		err = rows.Scan(&transaksi.Id, &transaksi.Tanggal_transaksi, &transaksi.Kasir_id)
		if err != nil {
			panic(err)
		}

		results = append(results, transaksi)
	}
	return
}

func InsertTransaksi(db *sql.DB, transaksi structs.Transaksi) (err error) {
	sql := "INSERT INTO transaksi (id, id_kasir, tanggal) VALUES ($1, $2, $3)"

	Tanggal_transaksi := time.Now()
	_, errs := db.Exec(sql, &transaksi.Id, &transaksi.Kasir_id, &Tanggal_transaksi)

	return errs
}

func GetDetailTransaksi(db *sql.DB, detail_transaksi structs.Detail_transaksi) (results []structs.Detail_transaksi, err error) {
	sql := `SELECT transaksi_id, menu_id, quantity, nama, m.harga 
	from detail_transaksi AS dt
	INNER JOIN menu as m ON m.id = dt.menu_id 
	WHERE dt.transaksi_id = $1`

	rows, err := db.Query(sql, detail_transaksi.Transaksi_id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		// detail_transaksi = structs.Detail_transaksi{}
		Total := detail_transaksi.Quantiti * detail_transaksi.Harga

		err = rows.Scan(&detail_transaksi.Transaksi_id, &detail_transaksi.Menu_id, &detail_transaksi.Quantiti, &detail_transaksi.Nama, &detail_transaksi.Harga)
		if err != nil {
			panic(err)
		}
		detail_transaksi.Total = Total
		results = append(results, detail_transaksi)
	}
	return
}

func InsertDetailTransaksi(db *sql.DB, Detail_transaksi structs.Detail_transaksi, menu structs.Menu) (err error) {
	sql := "INSERT INTO detail_transaksi (transaksi_id, menu_id, quantity) VALUES ($1, $2, $3)"

	_, errs := db.Exec(sql, &Detail_transaksi.Transaksi_id, &Detail_transaksi.Menu_id, &Detail_transaksi.Quantiti)

	return errs
}
