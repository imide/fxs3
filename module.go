package fxs3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/fx"
)

const ModuleName = "s3"

// FxRedis is the [Fx] redis module.
//
// [Fx]: https://github.com/uber-go/fx
var FxRedisModule = fx.Module(
	ModuleName,
	fx.Provide(
		fx.Annotate(NewDefaultS3ClientFactory, fx.As(new(S3ClientFactory))),
		NewFxS3Client,
	),
)

type FxS3ClientParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Options   s3.Options
	Factory   S3ClientFactory
}

func NewFxS3Client(p FxS3ClientParam) (*s3.Client, error) {
	client, err := p.Factory.Create()
	if err != nil {
		return nil, err
	}
	return client, nil
}
