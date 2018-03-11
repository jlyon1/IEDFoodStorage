package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/IEDFoodStorage/api"
	"github.com/jlyon1/IEDFoodStorage/config"
	"github.com/jlyon1/IEDFoodStorage/database"
	"net/http"
	"time"
)

func main() {

	cfg, err := config.New()

	if err != nil {
		fmt.Println("Failed to read config, exiting")
		panic("CFG error")
	}
	db := database.Database{}
	fmt.Printf("Food Manager Running, serving on port: %s\n", cfg.Port)

	fmt.Printf("Trying to connect to database: %s:%s\n", cfg.DBHostname, cfg.DBPort)

	err = db.Connect(cfg)
	if err != nil {
		fmt.Println("Failed to connect to database")
		for err != nil {
			<-time.After(5 * time.Second)
			fmt.Println("Retrying")

			err = db.Connect(cfg)
			if(err != nil){
				fmt.Println(err.Error())
			}
		}
		//panic(err.Error())
	}

	fmt.Println("Successfully connected")

	if !db.CheckLayout() {
		db.InitDatabase()
		db.CheckLayout()
	}

	r := mux.NewRouter()
	api := api.API{
		DB: &db,
	}

	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	r.HandleFunc("/inventory", api.IndexHandler).Methods("GET")
	r.HandleFunc("/remove/{id:[0-9]+}", api.RemoveFoodHandler).Methods("POST")

	http.ListenAndServe("0.0.0.0:"+cfg.Port, r)

}
