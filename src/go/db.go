package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func list() ([]StandardDeviation, error) {
	session, err := getDb()
	defer session.Close()
	var sds []StandardDeviation = nil
	if err == nil {
		err = session.DB(Config.DatabaseName).C("standardDeviation").Find(nil).All(&sds)
	}
	if sds == nil {
		sds = []StandardDeviation{}
	}
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
