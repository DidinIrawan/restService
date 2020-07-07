package models

type KaryawanReport struct {
	IdKaryawan     int    `json:"id_karyawan"`
	NamaLengkap    string `json:"nama_lengkap"`
	Alamat         string `json:"alamat"`
	Ttl            string `json:"ttl"`
	StatusKaryawan string `json:"status_karyawan"`
	Status         string `json:"status"`
	IdTunjangan    int    `json:"id_tunjangan"`
	NamaTunjangan  string `json:"nama_tunjangan"`
	IdGaji         int    `json:"id_gaji"`
	GajiPokok      int    `json:"gaji_pokok"`
}
