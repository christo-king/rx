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
	router := mux.NewRouter().StrictSlash(true)
	dir := http.Dir("/static/")
	fsrv := http.FileServer(dir)
	router.HandleFunc("/standardDeviation/{id}", HandleGetStandardDeviation).Methods("GET")
	router.Handle("/", fsrv)
	srv := http.Server{Addr: host, Handler: router}
	testdb := func(db *sql.DB) {
		db.Ping();
	}
	getDb(testdb)
	srv.ListenAndServe();
}

type DbResultHandler func(db *sql.DB)

// this is inefficient intentionally to avoid stateful issues
func getDb(h DbResultHandler) {
	// Create an sql.DB and check for errors
	db, err := sqlOpen("mysql", dburl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if ( h != nil ) {
		h(db);
	}
}
