package main

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var sqlOpen = sql.Open

const (
	dburl string = "test_user:test1@tcp(rxdb:3306)/testing"
	host string = "localhost:3001"
)

func main() {
	pingDb()

	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/static/")))

	list := HttpErrorHandler{HandleListStandardDeviations}
	router.HandleFunc("/standardDeviation", list.HandleHttpErrors).Methods("GET")
	get := HttpErrorHandler{HandleGetStandardDeviation}
	router.HandleFunc("/standardDeviation/{id}", get.HandleHttpErrors).Methods("GET")
	post := HttpErrorHandler{HandlePostStandardDeviation}
	router.HandleFunc("/standardDeviation", post.HandleHttpErrors).Methods("POST")

	srv := http.Server{Addr: host, Handler: router}
	srv.ListenAndServe();
}

func pingDb() {
	getDb(func(db *sql.DB) {
		db.Ping();
	})
}

// this is inefficient intentionally to avoid stateful issues
func getDb(dh func(db *sql.DB)) {
	// Create an sql.DB and check for errors
	db, err := sqlOpen("mysql", dburl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if ( dh != nil ) {
		dh(db);
	}
}
