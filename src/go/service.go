package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var Config = struct {
	Host         string
	DatabaseUrl  string
	DatabaseName string
}{
	Host:         os.Getenv("HOST"),
	DatabaseUrl:  os.Getenv("DB_MONGO"),
	DatabaseName: os.Getenv("DB_NAME"),
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	list := HttpErrorHandler{HandleListStandardDeviations}
	router.HandleFunc("/standardDeviation", list.HandleHttpErrors).Methods("GET")
	get := HttpErrorHandler{HandleGetStandardDeviation}
	router.HandleFunc("/standardDeviation/{id}", get.HandleHttpErrors).Methods("GET")
	post := HttpErrorHandler{HandlePostStandardDeviation}
	router.HandleFunc("/standardDeviation", post.HandleHttpErrors).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/static/")))

	srv := http.Server{Addr: Config.Host, Handler: router}
	srv.ListenAndServe()
}
