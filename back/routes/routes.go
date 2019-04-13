package routes

import (
	"github.com/gorilla/mux"
)

func newRoutes() *mux.Router {
	mux := mux.NewRouter()

	// redirect to frontend
	//mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	//mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	// a := mux.PathPrefix("/api").Subrouter()

	return mux
}
