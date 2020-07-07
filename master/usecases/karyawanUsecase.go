package usecases

import "restServiceApp/goWebMasterApi/master/models"

type KaryawanUsecase interface {
	GetAllKaryawan() ([]*models.Karyawan, error)
	GetAllKaryawanType() ([]*models.Karyawan, error)
	InsertKaryawan(karyawan models.Karyawan) error
	GetKaryawanById(id int) (*models.Karyawan, error)
	UpdateKaryawan(id int, karyawan models.Karyawan) error
	UpdateStatusKaryawan(id int, karyawan models.Karyawan) error
}
