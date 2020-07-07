package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"restServiceApp/goWebMasterApi/master/models"
	"restServiceApp/goWebMasterApi/master/usecases"
	"restServiceApp/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type KaryawanyHandler struct {
	karyawanUsecase usecases.KaryawanUsecase
}

func (h KaryawanyHandler) InsertKaryawan(w http.ResponseWriter, r *http.Request) {
	var karyawan models.Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = h.karyawanUsecase.InsertKaryawan(karyawan)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	byteOfKaryawan, err := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Insert", karyawan))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfKaryawan)
	log.Println("Insert successful")
}

func (h KaryawanyHandler) GetKaryawanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	karyawan, err := h.karyawanUsecase.GetKaryawanById(id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfKaryawan, err := json.Marshal(utils.GenerateResponse(http.StatusOK, "Data Berhasil Ditemukan", karyawan))
	if err != nil {
		w.Write([]byte("Oops, Something when wrong !!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfKaryawan)
}

func (h KaryawanyHandler) UpdateKaryawan(w http.ResponseWriter, r *http.Request) {
	var karyawan models.Karyawan
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)
	err := json.NewDecoder(r.Body).Decode(&karyawan)

	err = h.karyawanUsecase.UpdateKaryawan(id, karyawan)
	if err != nil {
		w.Write([]byte("Data not found !!"))
		log.Println(err)
		return
	}
	byteOfKaryawan, _ := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Update", karyawan))
	// w.Write([]byte("Update successful"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfKaryawan)
}

func (h KaryawanyHandler) UpdateStatusKaryawan(w http.ResponseWriter, r *http.Request) {
	var karyawan models.Karyawan
	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)
	_ = json.NewDecoder(r.Body).Decode(&karyawan)
	err := h.karyawanUsecase.UpdateKaryawan(id, karyawan)
	if err != nil {
		log.Println(err)
	}
	byteOfKaryawan, _ := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Update", karyawan))
	// w.Write([]byte("Update successful"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfKaryawan)
}

func (h KaryawanyHandler) GetAllKaryawan(w http.ResponseWriter, r *http.Request) {
	karyawans, err := h.karyawanUsecase.GetAllKaryawan()
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

func (h KaryawanyHandler) GetAllKaryawanType(w http.ResponseWriter, r *http.Request) {
	karyawans, err := h.karyawanUsecase.GetAllKaryawanType()
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

func KaryawanController(r *mux.Router, service usecases.KaryawanUsecase) {
	KaryawanHandler := KaryawanyHandler{service}
	r.HandleFunc("/karyawan", KaryawanHandler.InsertKaryawan).Methods(http.MethodPost)
	r.HandleFunc("/karyawan/{id}", KaryawanHandler.GetKaryawanById).Methods(http.MethodGet)
	r.HandleFunc("/karyawan/{id}", KaryawanHandler.UpdateKaryawan).Methods(http.MethodPut)
	r.HandleFunc("/karyawans", KaryawanHandler.GetAllKaryawan).Methods(http.MethodGet)
	r.HandleFunc("/karyawanstype", KaryawanHandler.GetAllKaryawanType).Methods(http.MethodGet)
}
