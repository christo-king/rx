package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	StandardDeviationPoints
	Id      bson.ObjectId `json:"id",bson:"_id,"`
	Created time.Time     `json:"created"`
	Answer  float64       `json:"answer"`
}
