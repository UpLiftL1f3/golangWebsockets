package main

import (
	"time"

	"github.com/UpLiftL1f3/tollCalc/types"
	"github.com/sirupsen/logrus"
)

type LogginMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogginMiddleware {
	return &LogginMiddleware{
		next: next,
	}
}

func (l *LogginMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
			"func":  "Producing to Kafka",
		}).Info()
	}(time.Now())
	return l.next.ProduceData(data)
}
