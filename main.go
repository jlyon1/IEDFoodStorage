package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlyon1/IEDFoodStorage/api"
	"github.com/jlyon1/IEDFoodStorage/config"
	"github.com/jlyon1/IEDFoodStorage/database"
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
			if err != nil {
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
	apiCFG := api.CFG{
		Enabled: true,
	}
	api := api.API{
		DB:  &db,
		Cfg: apiCFG,
	}

	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	r.HandleFunc("/get", api.GetHandler).Methods("GET")
	r.HandleFunc("/enabled", api.EnabledHandler).Methods("GET")
	r.HandleFunc("/enabled", api.UpdateStateHandler).Methods("Post")
	r.HandleFunc("/inventory", api.IndexHandler).Methods("GET")
	r.HandleFunc("/remove/{id:[0-9]+}", api.RemoveFoodHandler).Methods("POST")
	r.HandleFunc("/update", api.UpdateHandler).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe("0.0.0.0:"+cfg.Port, r)

}
