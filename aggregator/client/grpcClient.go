package client

import (
	"github.com/UpLiftL1f3/tollCalc/types"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	Endpoint string
	types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, nil)
	if err != nil {
		return nil, err
	}
	client := types.NewAggregatorClient(conn)

	return &GRPCClient{
		Endpoint:         endpoint,
		AggregatorClient: client,
	}, nil
}
