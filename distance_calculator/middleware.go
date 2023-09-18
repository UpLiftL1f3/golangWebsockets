package main

import (
	"log"
	"time"

	"github.com/UpLiftL1f3/tollCalc/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
			"func": "Calculate distance",
		}).Info()
	}(time.Now())
	dist, err = m.next.CalculateDistance(data)
	if err != nil {
		log.Fatal(err)
	}

	return dist, err
}
