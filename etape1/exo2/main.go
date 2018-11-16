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
	router.Methods("GET").Path("/contacts").HandlerFunc(getListeContacts)
	router.Methods("GET").Path("/contacts/{id}").HandlerFunc(getUnContact)
	router.Methods("POST").Path("/contacts/{id}").HandlerFunc(updateContact)
	router.Methods("PUT").Path("/contacts").HandlerFunc(createContact)
	router.Methods("DELETE").Path("/contacts/{id}").HandlerFunc(deleteContact)

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

func getListeContacts(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Liste des contact")
}

func getUnContact(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(req)
	idContact := params["id"]
	json.NewEncoder(w).Encode("Contact : " + idContact)
}

func updateContact(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	params := mux.Vars(req)
	idContact := params["id"]
	json.NewEncoder(w).Encode("Contact modifié " + idContact)
}

func createContact(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Contact créé")
}

func deleteContact(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(req)
	idContact := params["id"]
	json.NewEncoder(w).Encode("Contact supprimé : " + idContact)
}
