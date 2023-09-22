package client

import (
	"context"

	"github.com/UpLiftL1f3/tollCalc/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
