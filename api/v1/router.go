package v1

import (
	"github.com/gorilla/mux"
	"github.com/monofuel/japuraV2/api/v1/posts"
	"github.com/monofuel/japuraV2/api/v1/test"
)

//AddRoutes adds all v1 endpoints
func AddRoutes(r *mux.Router) {
	test.AddRoutes(r.PathPrefix("/test").Subrouter())
	posts.AddRoutes(r.PathPrefix("/posts").Subrouter())
}
