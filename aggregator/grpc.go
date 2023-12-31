package main

import (
	"context"

	"github.com/UpLiftL1f3/tollCalc/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

//* business layer -> business layer type (main type everyone needs to convert to)

//! Transport Layer
//- JSON -> types.Distance -> all done (same type)
//- GRPC -> types.AggregateRequest -> types.Distance
//- Webpack -> types.Webpack -> types.Distance

func (s *GRPCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
