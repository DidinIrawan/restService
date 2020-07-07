package repositories

import "restServiceApp/goWebMasterApi/master/models"

type ReportKaryawanRepository interface {
	GetTotalGajiPerbulan() ([]*models.ReportBulan, error)
	GetReportAllKaryawan() ([]*models.KaryawanReport, error)
	GetreportKaryawanById(id int) (*models.KaryawanReport, error)
}
