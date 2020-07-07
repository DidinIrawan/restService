package usecases

import (
	"log"
	"restServiceApp/goWebMasterApi/master/models"
	"restServiceApp/goWebMasterApi/master/repositories"
	"restServiceApp/utils"
)

type KaryawanUsecaseImpl struct {
	karyawanRepo repositories.KaryawanRepository
}

func (k KaryawanUsecaseImpl) GetAllKaryawanType() ([]*models.Karyawan, error) {
	karyawans, err := k.karyawanRepo.GetAllKaryawanType()
	if err != nil {
		return nil, err
	}
	return karyawans, nil
}

func (k KaryawanUsecaseImpl) GetAllKaryawan() ([]*models.Karyawan, error) {
	karyawans, err := k.karyawanRepo.GetAllKaryawan()
	if err != nil {
		return nil, err
	}
	return karyawans, nil
}

func (k KaryawanUsecaseImpl) UpdateStatusKaryawan(id int, karyawan models.Karyawan) error {
	panic("implement me")
}

func (k KaryawanUsecaseImpl) UpdateKaryawan(id int, karyawan models.Karyawan) error {
	_, err := k.karyawanRepo.GetKaryawanById(id)
	if err != nil {
		return err
	}
	err = k.karyawanRepo.UpdateKaryawan(id, karyawan)
	if err != nil {
		return err
	}
	return nil
}

func (k KaryawanUsecaseImpl) GetKaryawanById(id int) (*models.Karyawan, error) {

	karyawan, err := k.karyawanRepo.GetKaryawanById(id)
	if err != nil {
		return nil, err
	}
	return karyawan, nil
}

func (k KaryawanUsecaseImpl) InsertKaryawan(karyawan models.Karyawan) error {
	err := utils.DataValidation(karyawan.NamaLengkap, karyawan.Alamat, karyawan.Ttl, karyawan.StatusKaryawan)
	if err != nil {
		log.Print(err)
		return err
	}
	err = utils.ValidateInputLenCharacter(20, karyawan.NamaLengkap)
	if err != nil {
		return err
	}
	err = k.karyawanRepo.InsertKaryawann(karyawan)
	if err != nil {
		return err
	}
	return nil
}

func InitKaryawanUsecase(karyawan repositories.KaryawanRepository) KaryawanUsecase {
	return &KaryawanUsecaseImpl{karyawanRepo: karyawan}
}
