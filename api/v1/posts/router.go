package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/monofuel/japuraV2/db"
	"github.com/monofuel/japuraV2/util"
)

//AddRoutes adds the test endpoints
func AddRoutes(r *mux.Router) {
	r.Methods("GET").Path("/").Handler(util.AppHandler(getHandler))
	r.Methods("DELETE").Path("/{id}").Handler(util.AppHandler(deleteHandler))
	r.Methods("POST").Path("/").Handler(util.AppHandler(createHandler))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	database := db.GetDB()
	idParam := path.Base(r.URL.RequestURI())
	if idParam == "" {
		return &util.AppError{nil, "missing url ID", 400}
	}

	err := database.DeletePost(idParam)
	if err != nil {
		return &util.AppError{err, "failed to delete post", 500}
	}
	return nil
}

func getHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	database := db.GetDB()
	page, limit, err := util.GetPageLimitParameters(r)
	posts, err := database.GetPosts(page, limit)
	if err != nil {
		return &util.AppError{err, "failed to get posts", 500}
	}

	js, err := json.Marshal(posts)
	if err != nil {
		return util.NewAppError(err, "failed to marshal json", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
func createHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	database := db.GetDB()
	var post db.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return &util.AppError{err, "failed to parse body", 400}
	}
	id, err := database.AddPost(&post)
	if err != nil {
		return &util.AppError{err, "failed to create post", 500}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{id:%v}", id)
	return nil
}
