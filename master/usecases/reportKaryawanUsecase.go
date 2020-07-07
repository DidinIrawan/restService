package usecases

import "restServiceApp/goWebMasterApi/master/models"

type ReportKaryawanUsecase interface {
	GetTotalPerbulan() ([]*models.ReportBulan, error)
	GetReportAllKaryawan() ([]*models.KaryawanReport, error)
	GetReportKaryawanById(id int) (*models.KaryawanReport, error)
}
