package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kareemarab/space/api"
)

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world Endpoint.")
}

// NewRoutes ...
func NewRoutes(api *api.API) *mux.Router {
	mux := mux.NewRouter()

	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/hello", HelloWorld)

	return mux
}
