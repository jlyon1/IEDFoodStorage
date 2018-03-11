package api

import (
	"encoding/json"
	"github.com/jlyon1/IEDFoodStorage/database"
	"net/http"
)

type API struct {
	DB *database.Database
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 100; i++ {
		thing := api.DB.GetById(i)
		if thing.ID != 0 {
			WriteJSON(w, thing)
		} else {
			break
		}
	}
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
