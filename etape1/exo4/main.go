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
	router.Methods("GET").Path("/videos").HandlerFunc(getListeVideo)
	router.Methods("GET").Path("/videos/{id}").HandlerFunc(getUneVideo)
	router.Methods("POST").Path("/videos/{id}").HandlerFunc(updateVideo)
	router.Methods("PUT").Path("/videos").HandlerFunc(createVideo)
	router.Methods("DELETE").Path("/videos/{id}").HandlerFunc(deleteVideo)

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

func getListeVideo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Video"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Liste des vidéos"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)
}

func getUneVideo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(req)
	idContact := params["id"]

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Vidéo"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Vidéo : " + idContact

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)
}

func updateVideo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	params := mux.Vars(req)
	idContact := params["id"]
	meta.ObjectName = "Vidéo"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Vidéo modifiée " + idContact

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(reponse)

}

func createVideo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Vidéo"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Vidéo créés"

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reponse)

}

func deleteVideo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(req)
	idContact := params["id"]
	var meta metaReponse
	var reponse apiReponse

	meta.ObjectName = "Vidéo"
	meta.Count = 1
	meta.TotalCount = 1

	reponse.Meta = meta
	reponse.Data = "Vidéo supprimée : " + idContact

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reponse)

}
