package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/tx/send", SendTxHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	return r
}


