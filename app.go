package main

import (
	"restServiceApp/configs"
	"restServiceApp/goWebMasterApi/master"
)

func main() {
	db, _ := configs.ConnectionDB()
	router := configs.CreateRouter()
	master.Init(router, db)
	configs.RunServer(router)
}
