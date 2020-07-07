package usecases

import (
	"restServiceApp/goWebMasterApi/master/models"
	"restServiceApp/goWebMasterApi/master/repositories"
)

type ReportKaryawanUsecaseImpl struct {
	reportkaryawanRepo repositories.ReportKaryawanRepository
}

func (r ReportKaryawanUsecaseImpl) GetTotalPerbulan() ([]*models.ReportBulan, error) {
	karyawans, err := r.reportkaryawanRepo.GetTotalGajiPerbulan()
	if err != nil {
		return nil, err
	}
	return karyawans, nil
}

func (r ReportKaryawanUsecaseImpl) GetReportAllKaryawan() ([]*models.KaryawanReport, error) {
	karyawans, err := r.reportkaryawanRepo.GetReportAllKaryawan()
	if err != nil {
		return nil, err
	}
	return karyawans, nil
}

func (r ReportKaryawanUsecaseImpl) GetReportKaryawanById(id int) (*models.KaryawanReport, error) {
	karyawan, err := r.reportkaryawanRepo.GetreportKaryawanById(id)
	if err != nil {
		return nil, err
	}
	return karyawan, nil
}

func InitReportKaryawanUsecase(karyawan repositories.ReportKaryawanRepository) ReportKaryawanUsecase {
	return &ReportKaryawanUsecaseImpl{reportkaryawanRepo: karyawan}
}
