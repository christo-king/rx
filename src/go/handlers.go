package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	Id     int64
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

func HandlePostStandardDeviation(w http.ResponseWriter, r *http.Request) {
	body, ioerr := ioutil.ReadAll(r.Body)
	if ( ioerr != nil ) {
		log.Fatal(ioerr);
	}
	var sdp StandardDeviationPoints
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&sdp)
	if ( err != nil ) {
		log.Fatal(err)
	}
	sd := StandardDeviation{Id: -1, Answer : calcStdDev(sdp.Points), Points: sdp }
	stmt := "insert into standard_deviation_tbl(answer, input_data) values(?,?)";
	var saveNewStdDev = func(db *sql.DB) {
		res, dberr := db.Exec(stmt, sd.Answer, body);
		if ( dberr != nil ) {
			log.Fatal(dberr);
			w.WriteHeader(500)
		}
		newId, inserr := res.LastInsertId();
		if (inserr != nil) {
			log.Fatal(inserr)
			w.WriteHeader(500)
		}
		sd.Id = newId
		strout, marsherr := json.Marshal(sd)
		if ( marsherr != nil ) {
			log.Fatal(marsherr)
			w.WriteHeader(500)
		}
		w.Write(strout)
	}
	getDb(saveNewStdDev)
}