package models

type LaporanBelanja struct {
	ID         string `json:"id"`
	Tanggal    string `json:"tanggal"`
	NamaBarang string `json:"nama_barang"`
	Jumlah     int    `json:"jumlah"`
	Harga      int    `json:"harga"`
	Total      int    `json:"total"`
	Keterangan string `json:"keterangan"`
}
