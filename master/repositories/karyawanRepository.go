package repositories

import "restServiceApp/goWebMasterApi/master/models"

type KaryawanRepository interface {
	GetAllKaryawan() ([]*models.Karyawan, error)
	InsertKaryawann(karyawan models.Karyawan) error

	GetKaryawanById(id int) (*models.Karyawan, error)

	UpdateKaryawan(id int, karyawan models.Karyawan) error
	GetAllKaryawanType() ([]*models.Karyawan, error)
}
