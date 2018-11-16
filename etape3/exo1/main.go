package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// http://www.gorillatoolkit.org/pkg/mux

func main() {
	var router *mux.Router
	var err error

	router = initialisationRoutes()

	router.Methods("GET").Path("/").HandlerFunc(helloErnesto)
	err = http.ListenAndServe(":8880", router)
	if err != nil {
		log.Fatal(err)
	}
}

func initialisationRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

func helloErnesto(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Hello Ernesto")
}
