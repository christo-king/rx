package main

import (
	"time"
)

type StandardDeviationPoints struct {
	Points []float64 `json:"points"`
}

type StandardDeviation struct {
	StandardDeviationPoints
	Id      string    `json:"id" bson:"_id"`
	Created time.Time `json:"created"`
	Answer  float64   `json:"answer"`
}
