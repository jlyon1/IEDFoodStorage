package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/IEDFoodStorage/api"
	"github.com/jlyon1/IEDFoodStorage/config"
	"github.com/jlyon1/IEDFoodStorage/database"

	"net/http"
)

func main() {

	cfg, err := config.New()

	if err != nil {
		fmt.Println("Failed to read config, exiting")
		panic("CFG error")
	}
	db := database.Database{}
	fmt.Printf("Food Manager Running, serving on port: %s\n", cfg.Port)

	fmt.Println("Trying to connect to database")

	err = db.Connect(cfg)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err.Error())
	}

	fmt.Println("Successfully connected")

	if !db.CheckLayout() {
		db.InitDatabase()
		db.CheckLayout()
	}
	db.GetById(1)
	r := mux.NewRouter()
	api := api.API{
		DB: &db,
	}

	r.HandleFunc("/", api.IndexHandler).Methods("GET")

	http.ListenAndServe("0.0.0.0:"+cfg.Port, r)

}
