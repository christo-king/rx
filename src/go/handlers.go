package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"encoding/json"
	"io/ioutil"
	"errors"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	StandardDeviationPoints
	Id     int64 `json:"id"`
	Answer float64 `json:"answer"`
}

func HandleListStandardDeviations(w http.ResponseWriter, r *http.Request) HttpError {
	stmt := "select sd.id, sd.answer, sd.input_data from standard_deviation_tbl"
	listerr := HttpOK()
	var unmarshallStdDev = func(db *sql.DB) {
		rows, qerr := db.Query(stmt)
		if ( qerr != nil ) {
			listerr = HttpResponse{500, errors.New("Unable to query for Standard Deviations"), qerr}
			return;
		}
		for rows.Next() {
			var sd StandardDeviation
			var sdpStr string
			err := rows.Scan(&sd.Id, &sd.Answer, &sdpStr)
			switch {
			case err == sql.ErrNoRows:
				w.Write([]byte("{}"))
				return;
			case err != nil:
				listerr = NewHttpError(500, "Error selecting specified Standard Deviation")
				return;
			default:
				jsonerr := json.Unmarshal([]byte(sdpStr), &sd)
				if ( jsonerr != nil) {
					listerr = NewHttpError(500, "Error selecting specified Standard Deviation")
					return;
				}
				strout, jsonerr := json.Marshal(sd);
				if ( jsonerr != nil ) {
					listerr = NewLogHttpError(500, "Invalid standard deviation result", jsonerr)
					return;
				}
				w.Write(strout);
			}
		}
	}
	getDb(unmarshallStdDev)
	return listerr
}

func HandleGetStandardDeviation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stmt := "select id, answer, input_data from standard_deviation_tbl where id=?"
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
			jsonerr := json.Unmarshal([]byte(sdpStr), &sd)
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
	var bodybytes, strerr = ioutil.ReadAll(r.Body)
	var bodystr = string(bodybytes)
	if ( strerr != nil ) {
		log.Fatal(strerr)
		w.WriteHeader(500)
	}
	var sd StandardDeviation
	sd.Id = -1
	jerr := json.Unmarshal(bodybytes, &sd)
	if ( jerr != nil ) {
		log.Fatal(jerr)
		w.WriteHeader(500)
	}
	sd.Answer = calcStdDev(sd.Points)

	stmt := "insert into standard_deviation_tbl(answer, input_data) values(?,?)";
	var saveNewStdDev = func(db *sql.DB) {
		res, dberr := db.Exec(stmt, sd.Answer, bodystr);
		if ( dberr != nil ) {
			log.Fatal(dberr);
			w.WriteHeader(500)
		}
		newId, inserr := res.LastInsertId();
		if (inserr != nil) {
			log.Fatal(inserr)
			http.Error(w, "", 500)
		}
		sd.Id = newId
		json.NewEncoder(w).Encode(sd)
	}
	getDb(saveNewStdDev)
}
