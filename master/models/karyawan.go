package models

type Karyawan struct {
	IdKaryawan     int    `json:"id_karyawan"`
	NamaLengkap    string `json:"nama_lengkap"`
	Alamat         string `json:"alamat"`
	Ttl            string `json:"ttl"`
	StatusKaryawan string `json:"status_karyawan"`
	Status         string `json:"status"`
}
