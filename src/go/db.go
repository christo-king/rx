package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const STANDARD_DEVIATION = "standardDeviation"

func list() ([]StandardDeviation, error) {
	session, err := getDb()
	defer session.Close()
	var sds []StandardDeviation = nil
	if err == nil {
		err = session.DB(Config.DatabaseName).C(STANDARD_DEVIATION).Find(nil).All(&sds)
	}
	if sds == nil {
		sds = []StandardDeviation{}
	}
	return sds, err
}

func get(id string) (StandardDeviation, error) {
	session, err := getDb()
	defer session.Close()
	var sd StandardDeviation
	if err == nil {
		err = session.DB(Config.DatabaseName).C(STANDARD_DEVIATION).FindId(bson.M{"_id":bson.ObjectIdHex(id)}).One(&sd)
	}
	log.Printf("CLOBBO %+v", err)
	return sd, err
}

func save(sd *StandardDeviation) (bool, error) {
	success := false
	session, err := getDb()
	if err == nil {
		sd.Id = bson.NewObjectId().Hex()
		session.DB(Config.DatabaseName).C(STANDARD_DEVIATION).Insert(sd)
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
