package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/IEDFoodStorage/database"
	"github.com/jlyon1/IEDFoodStorage/model"
	"net/http"
	"strconv"
)

type API struct {
	DB *database.Database
}

func (api *API) GetHandler(w http.ResponseWriter, r *http.Request) {
	f := api.DB.GetAll()
	WriteJSON(w, f)

}

func (api *API) RemoveFoodHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		http.Error(w, "Invalid format", 500)
		return
	}

	if api.DB.Remove(id) {
		w.Write([]byte("removed"))
		fmt.Printf("Removed %d\n", id)
	} else {
		http.Error(w, "Failed", 500)
	}

}

func (api *API) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	food := model.Food{}
	err := json.NewDecoder(r.Body).Decode(&food)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		http.Error(w, "Invalid format", 500)
		return
	}
	if !api.DB.Update(food) {
		http.Error(w, "Error updating db", 500)
		return
	}
	w.Write([]byte("updated"))

}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
