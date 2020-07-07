package repositories

import (
	"database/sql"
	"restServiceApp/goWebMasterApi/master/models"
)

type ReportKaryawanRepoImpl struct {
	db *sql.DB
}

func (r ReportKaryawanRepoImpl) GetTotalGajiPerbulan() ([]*models.ReportBulan, error) {
	var karyawans []*models.ReportBulan
	query := "SELECT month(g.tgl),sum(g.gapok) as Perbulan FROM detail_karyawan dk JOIN m_karyawan k ON dk.id_karyawan = k.id_karyawan JOIN m_tunjangan t ON dk.id_tunjangan = t.id_tunjangan JOIN m_gaji g ON dk.id_gaji = g.id_gaji where k.status like '%A%' group by 1;"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		karyawan := models.ReportBulan{}
		err := rows.Scan(&karyawan.Bulan, &karyawan.Total)
		if err != nil {
			return nil, err
		}
		karyawans = append(karyawans, &karyawan)
	}
	return karyawans, nil
}

func (r ReportKaryawanRepoImpl) GetReportAllKaryawan() ([]*models.KaryawanReport, error) {
	var karyawans []*models.KaryawanReport
	query := "SELECT dk.id_karyawan, k.nama_lengkap, k.alamat, k.ttl, k.statusKaryawan,k.status,t.id_tunjangan,t.nama_tunjangan,g.id_gaji,g.gapok FROM detail_karyawan dk JOIN m_karyawan k ON dk.id_karyawan = k.id_karyawan JOIN m_tunjangan t ON dk.id_tunjangan = t.id_tunjangan JOIN m_gaji g ON dk.id_gaji = g.id_gaji;"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		karyawan := models.KaryawanReport{}
		err := rows.Scan(&karyawan.IdKaryawan, &karyawan.NamaLengkap, &karyawan.Alamat, &karyawan.Ttl, &karyawan.StatusKaryawan, &karyawan.Status, &karyawan.IdTunjangan, &karyawan.NamaTunjangan, &karyawan.IdGaji, &karyawan.GajiPokok)
		if err != nil {
			return nil, err
		}
		karyawans = append(karyawans, &karyawan)
	}
	return karyawans, nil
}

func (r ReportKaryawanRepoImpl) GetreportKaryawanById(id int) (*models.KaryawanReport, error) {
	karyawan := new(models.KaryawanReport)
	query := "SELECT dk.id_karyawan, k.nama_lengkap, k.alamat, k.ttl, k.statusKaryawan,k.status,t.id_tunjangan,t.nama_tunjangan,g.id_gaji,g.gapok FROM detail_karyawan dk JOIN m_karyawan k ON dk.id_karyawan = k.id_karyawan JOIN m_tunjangan t ON dk.id_tunjangan = t.id_tunjangan JOIN m_gaji g ON dk.id_gaji = g.id_gaji where dk.id_karyawan=?;"
	if err := r.db.QueryRow(query, id).Scan(&karyawan.IdKaryawan, &karyawan.NamaLengkap, &karyawan.Alamat, &karyawan.Ttl, &karyawan.StatusKaryawan, &karyawan.Status, &karyawan.IdTunjangan, &karyawan.NamaTunjangan, &karyawan.IdGaji, &karyawan.GajiPokok); err != nil {
		return nil, err
	}
	return karyawan, nil
}

func InitReportKaryawanRepoImpl(db *sql.DB) ReportKaryawanRepository {
	return &ReportKaryawanRepoImpl{db}
}
