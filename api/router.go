package api

import (
	"github.com/gorilla/mux"
	"github.com/monofuel/japuraV2/api/v1"
)

//AddRoutes adds all api endpoints
func AddRoutes(r *mux.Router) {
	v1.AddRoutes(r.PathPrefix("/v1").Subrouter())
}
