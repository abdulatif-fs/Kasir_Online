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
