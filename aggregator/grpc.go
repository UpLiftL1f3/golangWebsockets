package main

import "github.com/UpLiftL1f3/tollCalc/types"

type GRPCServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCServer(svc Aggregator) *GRPCServer {
	return &GRPCServer{
		svc: svc,
	}
}

//* business layer -> business layer type (main type everyone needs to convert to)

//! Transport Layer
//- JSON -> types.Distance -> all done (same type)
//- GRPC -> types.AggregateRequest -> types.Distance
//- Webpack -> types.Webpack -> types.Distance

func (s *GRPCServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
