package main

import (
	"demo/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/websites", service.Getdata).Methods(http.MethodGet)
	r.HandleFunc("/websites", service.CreateData).Methods(http.MethodPost)
	r.HandleFunc("/website", service.CheckQuery).Methods(http.MethodGet)

	http.ListenAndServe("localhost:3000", r)
}
