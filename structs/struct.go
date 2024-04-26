package structs

type Menu struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

type Kasir struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Transaksi struct {
	Id                int         `json:"id"`
	Kasir_id          int         `json:"kasir_id"`
	Tanggal_transaksi interface{} `json:"tanggal_transaksi"`
}

type Detail_transaksi struct {
	Transaksi_id int    `json:"transaksi_id"`
	Menu_id      int    `json:"menu_id"`
	Quantiti     int    `json:"quantiti"`
	Nama         string `json:"nama"`
	Harga        int    `json:"harga"`
	Total        int    `json:"total"`
}
