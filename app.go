package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllCountriesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindCountryEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateCountryEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateCountryEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteCountryEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/countries", AllCountriesEndPoint).Methods("GET")
	r.HandleFunc("/countries", CreateCountryEndPoint).Methods("POST")
	r.HandleFunc("/countries", UpdateCountryEndPoint).Methods("PUT")
	r.HandleFunc("/countries", DeleteCountryEndPoint).Methods("DELETE")
	r.HandleFunc("/countries/{id}", FindCountryEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}