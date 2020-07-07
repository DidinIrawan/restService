package repositories

import (
	"database/sql"
	"restServiceApp/goWebMasterApi/master/models"
)

type KaryawanRepoImpl struct {
	db *sql.DB
}

func (k KaryawanRepoImpl) GetAllKaryawanType() ([]*models.Karyawan, error) {
	var karyawans []*models.Karyawan
	query := "SELECT id_karyawan,nama_lengkap,alamat, ttl,statusKaryawan,status from m_karyawan where statusKaryawan like '%permanen%' and status like '%A%'"
	rows, err := k.db.Query(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		karyawan := models.Karyawan{}
		err := rows.Scan(&karyawan.IdKaryawan, &karyawan.NamaLengkap, &karyawan.Alamat, &karyawan.Ttl, &karyawan.StatusKaryawan, &karyawan.Status)
		if err != nil {
			return nil, err
		}
		karyawans = append(karyawans, &karyawan)
	}
	return karyawans, nil
}

func (k KaryawanRepoImpl) GetAllKaryawan() ([]*models.Karyawan, error) {
	var karyawans []*models.Karyawan
	query := "SELECT id_karyawan,nama_lengkap,alamat, ttl,statusKaryawan,status from m_karyawan"
	rows, err := k.db.Query(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		karyawan := models.Karyawan{}
		err := rows.Scan(&karyawan.IdKaryawan, &karyawan.NamaLengkap, &karyawan.Alamat, &karyawan.Ttl, &karyawan.StatusKaryawan, &karyawan.Status)
		if err != nil {
			return nil, err
		}
		karyawans = append(karyawans, &karyawan)
	}
	return karyawans, nil
}

func (k KaryawanRepoImpl) UpdateKaryawan(id int, karyawan models.Karyawan) error {
	tx, err := k.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE m_karyawan SET nama_lengkap=?, alamat=?,ttl=?,statusKaryawan=? WHERE id_karyawan = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(karyawan.NamaLengkap, karyawan.Alamat, karyawan.Ttl, karyawan.StatusKaryawan, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (k KaryawanRepoImpl) GetKaryawanById(id int) (*models.Karyawan, error) {
	karyawan := new(models.Karyawan)
	query := "SELECT id_karyawan,nama_lengkap, alamat,ttl,statusKaryawan,status from m_karyawan where id_karyawan=?"
	if err := k.db.QueryRow(query, id).Scan(&karyawan.IdKaryawan, &karyawan.NamaLengkap, &karyawan.Alamat, &karyawan.Ttl, &karyawan.StatusKaryawan, &karyawan.Status); err != nil {
		return nil, err
	}
	return karyawan, nil
}

func (k KaryawanRepoImpl) InsertKaryawann(karyawan models.Karyawan) error {
	tx, err := k.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO m_karyawan (nama_lengkap,alamat,ttl,statusKaryawan) VALUES (?,?,?,?)"
	smt, err := tx.Exec(query, karyawan.NamaLengkap, karyawan.Alamat, karyawan.Ttl, karyawan.StatusKaryawan)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = smt.LastInsertId()
	tx.Commit()
	return nil
}

func InitKaryawanRepoImpl(db *sql.DB) KaryawanRepository {
	return &KaryawanRepoImpl{db}
}
