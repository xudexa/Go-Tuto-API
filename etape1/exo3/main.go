package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// http://www.gorillatoolkit.org/pkg/mux

type metaReponse struct {
	ObjectName string `json:"ObjectName"`
	TotalCount int    `json:"TotalCount"`
	Offset     int    `json:"Offset"`
	Count      int    `json:"Count"`
}

type apiReponse struct {
	Meta metaReponse `json:"Meta"`
	Data string      `json:"Data"`
}

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode("Hello Ernesto")
}

func getListeContacts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Contact"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Liste des contact"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)
}

func getUnContact(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(req)
	idContact := params["id"]

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Contact"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Contact : " + idContact

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)
}

func updateContact(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	params := mux.Vars(req)
	idContact := params["id"]
	meta.ObjectName = "Contact"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Contact modifié " + idContact

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(reponse)

}

func createContact(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Contact"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Contact créé"

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reponse)

}

func deleteContact(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(req)
	idContact := params["id"]
	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Contact"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Contact supprimé : " + idContact

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)

}
