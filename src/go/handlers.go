package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	StandardDeviationPoints
	Id     int64   `json:"id"`
	Answer float64 `json:"answer"`
}

func HandleListStandardDeviations(w http.ResponseWriter, r *http.Request) HttpError {
	listerr := HttpOK()
	var sdlist []StandardDeviation = []StandardDeviation{}

	if listerr.code != 200 {
		return listerr
	}
	strout, jsonerr := json.Marshal(sdlist)
	w.Write(strout)
	if ( jsonerr != nil ) {
		listerr = NewLogHttpError(500, "Invalid standard deviation result", jsonerr)
	} else {
		w.Write(strout);
	}
	return listerr
}

func HandleGetStandardDeviation(w http.ResponseWriter, r *http.Request) HttpError {
	//vars := mux.Vars(r)
	geterr := HttpOK()
	strout, jsonerr := json.Marshal([]StandardDeviation{});
	if (jsonerr != nil) {
		geterr = NewLogHttpError(500, "Unable to serialize standard deviation", jsonerr)
	} else {
		w.Write(strout)
	}
	log.Print(strout)
	w.Write(strout);
	return geterr;
}

func HandlePostStandardDeviation(w http.ResponseWriter, r *http.Request) HttpError {
	posterr := HttpOK()

	var bodybytes, strerr = ioutil.ReadAll(r.Body)
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
	sd.Answer = calcStdDev(sd.Points)

	json.NewEncoder(w).Encode(sd)

	return posterr
}
