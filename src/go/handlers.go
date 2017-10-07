package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	StandardDeviationPoints
	Id     bson.ObjectId `json:"id",bson:"_id,"`
	Answer float64       `json:"answer"`
}

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

func list() ([]StandardDeviation, error) {
	session, err := getDb()
	var sds []StandardDeviation = nil
	if err == nil {
		err = session.DB(Config.DatabaseName).C("standardDeviation").Find(nil).All(&sds)
	}
	defer session.Close()
	return sds, err
}
func save(sd *StandardDeviation) (bool, error) {
	success := false
	session, err := getDb()
	if err == nil {
		id := bson.NewObjectId()
		sd.Id = id
		session.DB(Config.DatabaseName).C("standardDeviation").Insert(sd)
		success = true
	}
	defer session.Close()
	return success, nil
}

func getDb() (*mgo.Session, error) {
	sess, err := mgo.Dial(Config.DatabaseUrl)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return sess, nil
}
