package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
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
	stmt := "select sd.id, sd.answer, sd.input_data from standard_deviation_tbl sd"
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
				return
			case err != nil:
				listerr = NewHttpError(500, "Error selecting specified Standard Deviation")
				return
			default:
				jsonerr := json.Unmarshal([]byte(sdpStr), &sd)
				if ( jsonerr != nil) {
					listerr = NewHttpError(500, "Error selecting specified Standard Deviation")
					return
				}
				strout, jsonerr := json.Marshal(sd);
				if ( jsonerr != nil ) {
					listerr = NewLogHttpError(500, "Invalid standard deviation result", jsonerr)
					return
				}
				w.Write(strout);
			}
		}
	}
	getDb(unmarshallStdDev)
	return listerr
}

func HandleGetStandardDeviation(w http.ResponseWriter, r *http.Request) HttpError {
	vars := mux.Vars(r)
	stmt := "select id, answer, input_data from standard_deviation_tbl where id=?"
	geterr := HttpOK()
	var unmarshallStdDev = func(db *sql.DB) {
		var sd StandardDeviation
		var sdpStr string
		err := db.QueryRow(stmt, vars["id"]).Scan(&sd.Id, &sd.Answer, &sdpStr)
		switch {
		case err == sql.ErrNoRows:
			geterr = NewHttpError(404, "Standard deviation Not found " + string(vars["id"]))
			return
		case err != nil:
			geterr = NewLogHttpError(500, "Unable to query database", err)
			return
		default:
			jsonerr := json.Unmarshal([]byte(sdpStr), &sd)
			if ( jsonerr != nil) {
				geterr = NewLogHttpError(500, "database values for this standard deviation are invalid", jsonerr)
				return
			}
			strout, jsonerr := json.Marshal(sd);
			if ( jsonerr != nil ) {
				geterr = NewLogHttpError(500, "database values for this standard deviation are invalid", jsonerr)
				return
			}
			w.Write(strout);
		}
	}
	getDb(unmarshallStdDev)
	return geterr;
}

func HandlePostStandardDeviation(w http.ResponseWriter, r *http.Request) HttpError {
	posterr := HttpOK()

	var bodybytes, strerr = ioutil.ReadAll(r.Body)
	var bodystr = string(bodybytes)
	if ( strerr != nil ) {
		posterr = NewLogHttpError(400, "Invalid standard deviation body", strerr)
		return posterr
	}

	var sd StandardDeviation
	jerr := json.Unmarshal(bodybytes, &sd)
	if ( jerr != nil ) {
		posterr = NewLogHttpError(400, "Unable to decode standard deviation", jerr)
		return posterr
	}
	sd.Id = -1
	sd.Answer = calcStdDev(sd.Points)

	stmt := "insert into standard_deviation_tbl(answer, input_data) values(?,?)";
	var saveNewStdDev = func(db *sql.DB) {
		res, dberr := db.Exec(stmt, sd.Answer, bodystr);
		if ( dberr != nil ) {
			posterr = NewLogHttpError(500, "Unable to insert new standard deviation", dberr)
			return
		}
		newId, inserr := res.LastInsertId();
		if (inserr != nil) {
			posterr = NewLogHttpError(500, "Unable to get id for new standard deviation", dberr)
			return
		}
		sd.Id = newId
		json.NewEncoder(w).Encode(sd)
	}
	getDb(saveNewStdDev)
	return posterr
}
