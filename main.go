package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/monofuel/japuraV2/api"
	"github.com/monofuel/japuraV2/auth"
	"github.com/monofuel/japuraV2/util"
)

var webPort = os.Getenv("PORT")
var templates = template.Must(template.ParseFiles("public/index.html"))

func init() {
	if webPort == "" {
		webPort = "8085"
	}

	//TODO connect to database
}

func main() {
	fmt.Println("starting example webapp")
	registerHandlers()

	log.Printf("listening on port %s", webPort)
	log.Fatal(http.ListenAndServe(":"+webPort, nil))
}

func registerHandlers() {

	r := mux.NewRouter()

	r.Methods("GET").Path("/").Handler(util.AppHandler(rootHandler))

	jsFs := http.FileServer(http.Dir("public/js"))
	imageFs := http.FileServer(http.Dir("public/images"))
	cssFs := http.FileServer(http.Dir("public/css"))
	certFs := http.FileServer(http.Dir("public/.well-known"))
	r.Methods("GET").PathPrefix("/js/").Handler(http.StripPrefix("/js", jsFs))
	r.Methods("GET").PathPrefix("/images/").Handler(http.StripPrefix("/images", imageFs))
	r.Methods("GET").PathPrefix("/css/").Handler(http.StripPrefix("/css", cssFs))
	r.Methods("GET").PathPrefix("/.well-known/").Handler(http.StripPrefix("/.well-known/", certFs))

	api.AddRoutes(r)
	auth.AddRoutes(r)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) *util.AppError {
	user, _, err := auth.AuthUser(w, r)
	var pageSettings struct {
		Admin       bool
		DisplayName string
	}
	if err != nil {
		fmt.Printf("failed to auth user: %v", err)
	} else {
		pageSettings.Admin = user.Admin
		pageSettings.DisplayName = user.Username
	}
	err = templates.ExecuteTemplate(w, "index.html", pageSettings)
	if err != nil {
		return util.NewAppError(err, "failed to load index.html", 500)
	}
	return nil
}
