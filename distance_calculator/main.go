package main

import (
	"log"

	"github.com/UpLiftL1f3/tollCalc/aggregator/client"
)

const (
	kafkaTopic         = "obuData"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

// * Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport
func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	// httpClient := client.NewHTTPClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}
