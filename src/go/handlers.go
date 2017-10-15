package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func HandleListStandardDeviations(w http.ResponseWriter, r *http.Request) HttpError {
	listerr := HttpOK()

	sdlist, err := list()

	if err != nil {
		return NewLogHttpError(500, "Unable to list standard deviations", err)
	}
	log.Println(fmt.Sprintf("DEVS: %v", sdlist))
	json.NewEncoder(w).Encode(sdlist)
	return listerr
}

func HandleGetStandardDeviation(w http.ResponseWriter, r *http.Request) HttpError {
	vars := mux.Vars(r)
	sd, err := get(vars["id"])
	if err != nil {
		return NewLogHttpError(500, "Unable to serialize standard deviation", err)
	}
	json.NewEncoder(w).Encode(sd)
	return HttpOK()
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
	sd.Created = time.Now()
	success, err := save(&sd)
	log.Println(fmt.Sprintf(" something %+v", sd.Id))

	if success {
		json.NewEncoder(w).Encode(sd)
	} else {
		posterr = NewLogHttpError(500, "Unable to save standard deviation", err)
	}

	return posterr
}
