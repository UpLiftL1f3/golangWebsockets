package main

import (
	"context"
	"log"
	"time"

	"github.com/UpLiftL1f3/tollCalc/aggregator/client"
	"github.com/UpLiftL1f3/tollCalc/types"
)

func main() {
	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 123.12,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}

}
