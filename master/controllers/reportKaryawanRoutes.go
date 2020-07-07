package controllers

import (
	"encoding/json"
	"net/http"
	"restServiceApp/goWebMasterApi/master/usecases"
	"restServiceApp/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type ReportKaryawanyHandler struct {
	ReportkaryawanUsecase usecases.ReportKaryawanUsecase
}

func (h ReportKaryawanyHandler) GetReportKaryawanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)
	karyawan, err := h.ReportkaryawanUsecase.GetReportKaryawanById(id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfReportKaryawan, err := json.Marshal(utils.GenerateResponse(http.StatusOK, "Data Berhasil Ditemukan", karyawan))
	if err != nil {
		w.Write([]byte("Oops, Something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfReportKaryawan)
}

func (h ReportKaryawanyHandler) GetReportAllKaryawan(w http.ResponseWriter, s *http.Request) {
	karyawans, err := h.ReportkaryawanUsecase.GetReportAllKaryawan()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfkaryawans, err := json.Marshal(utils.GenerateResponse(http.StatusOK, "Data Berhasil Ditampilkan", karyawans))
	if err != nil {
		w.Write([]byte("Oops, Something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfkaryawans)
}

func (h ReportKaryawanyHandler) GetReportTotalPerbulan(w http.ResponseWriter, r *http.Request) {
	totalPerbulan, err := h.ReportkaryawanUsecase.GetTotalPerbulan()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfTotalPerbulan, err := json.Marshal(utils.GenerateResponse(http.StatusOK, "Data Berhasil Ditampilkan", totalPerbulan))
	if err != nil {
		w.Write([]byte("Oops, Something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfTotalPerbulan)
}
func ReportKaryawanController(r *mux.Router, service usecases.ReportKaryawanUsecase) {
	ReportKaryawanHandler := ReportKaryawanyHandler{service}
	r.HandleFunc("/reportkaryawan/{id}", ReportKaryawanHandler.GetReportKaryawanById).Methods(http.MethodGet)
	r.HandleFunc("/reportkaryawans", ReportKaryawanHandler.GetReportAllKaryawan).Methods(http.MethodGet)
	r.HandleFunc("/reporttotalperbulan", ReportKaryawanHandler.GetReportTotalPerbulan).Methods(http.MethodGet)
}
