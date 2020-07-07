package master

import (
	"database/sql"
	"restServiceApp/goWebMasterApi/master/controllers"
	"restServiceApp/goWebMasterApi/master/repositories"
	"restServiceApp/goWebMasterApi/master/usecases"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	//karyawanRepo
	karyawanRepo := repositories.InitKaryawanRepoImpl(db)
	karyawanUsecase := usecases.InitKaryawanUsecase(karyawanRepo)
	controllers.KaryawanController(r, karyawanUsecase)

	//ReportkaryawanRepo
	ReportkaryawanRepo := repositories.InitReportKaryawanRepoImpl(db)
	ReportkaryawanUsecase := usecases.InitReportKaryawanUsecase(ReportkaryawanRepo)
	controllers.ReportKaryawanController(r, ReportkaryawanUsecase)
}
