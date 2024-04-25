package structs

type Menu struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

type Kasir struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type Transaksi struct {
	Id                int         `json:"id"`
	Kasir_id          int         `json:"kasir_id"`
	Tanggal_transaksi interface{} `json:"tanggal_transaksi"`
	Total_belanja     int
}

type Detail_transaksi struct {
	Transaksi_id int `json:"transaksi_id"`
	Menu_id      int `json:"menu_id"`
	Quantiti     int `json:"quantiti"`
	Total        int
}
