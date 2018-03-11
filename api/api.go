package api

import (
	"github.com/jlyon1/IEDFoodStorage/database"
	"net/http"
)

type API struct {
	DB *database.Database
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
