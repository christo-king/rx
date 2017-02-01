package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"encoding/json"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	Id     int
	Answer float64
	Points StandardDeviationPoints
}

func HandleGetStandardDeviation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stmt := "select sd.id, sd.answer, sd.input_data from standard_deviation_tbl sd where sd.id=?"
	var unmarshallStdDev = func(db *sql.DB) {
		var sd StandardDeviation
		var sdpStr string
		err := db.QueryRow(stmt, vars["id"]).Scan(&sd.Id, &sd.Answer, &sdpStr)
		switch {
		case err == sql.ErrNoRows:
			w.WriteHeader(404)
		case err != nil:
			log.Fatal(err)
		default:
			jsonerr := json.Unmarshal([]byte(sdpStr), &sd.Points)
			if ( jsonerr != nil) {
				panic(jsonerr)
			}
			strout, jsonerr := json.Marshal(sd);
			if ( jsonerr != nil ) {
				panic(jsonerr);
			}
			w.Write(strout);
		}

	}
	getDb(unmarshallStdDev)
}

