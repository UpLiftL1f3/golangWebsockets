package main

import (
	"time"

	"github.com/UpLiftL1f3/tollCalc/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(data types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"func":  "AggregateDistance",
			"took":  time.Since(start),
			"error": err,
		}).Info()
	}(time.Now())

	err = m.next.AggregateDistance(data)
	return
}

func (m *LogMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"func":     "Calculate Invoice",
			"took":     time.Since(start),
			"error":    err,
			"obuID":    obuID,
			"amount":   amount,
			"distance": distance,
		}).Info()
	}(time.Now())

	inv, err = m.next.CalculateInvoice(obuID)
	return
}
