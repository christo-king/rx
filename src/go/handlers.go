package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

func HandleListStandardDeviations(w http.ResponseWriter, r *http.Request) HttpError {
	listerr := HttpOK()

	sdlist, err := list()

	if err != nil {
		return NewLogHttpError(500, "Unable to list standard deviations", err)
	}
	json.NewEncoder(w).Encode(sdlist)
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

	var bodybytes, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return NewLogHttpError(400, "Invalid standard deviation body", err)
	}

	var sd StandardDeviation
	err = json.Unmarshal(bodybytes, &sd)
	if err != nil {
		return NewLogHttpError(400, "Unable to decode standard deviation", err)
	}
	sd.Answer = calcStdDev(sd.Points)
	success, err := save(&sd)
	log.Println(fmt.Sprintf(" something %+v", sd.Id))

	if success {
		json.NewEncoder(w).Encode(sd)
	} else {
		posterr = NewLogHttpError(500, "Unable to save standard deviation", err)
	}

	return posterr
}
