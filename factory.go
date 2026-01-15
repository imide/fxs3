package fxs3

import (
	"context"

	"github.com/ankorstore/yokai/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ S3ClientFactory = (*DefaultS3ClientFactory)(nil)

type S3ClientFactory interface {
	Create() (*s3.Client, error)
}

type DefaultS3ClientFactory struct {
	config *config.Config
	ctx    context.Context
}

func NewDefaultS3ClientFactory(config *config.Config) S3ClientFactory {
	return &DefaultS3ClientFactory{
		config: config,
	}
}

func (f *DefaultS3ClientFactory) Create() (*s3.Client, error) {
	sdkConfig, err := awsConfig.LoadDefaultConfig(f.ctx, awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		f.config.GetString("modules.s3.aws_access_key_id"),
		f.config.GetString("aws_secret_access_key"),
		"",
	),
	))

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(f.config.GetString("modules.s3.endpoint"))
		o.Region = "auto"
		o.UsePathStyle = false
		o.DisableLogOutputChecksumValidationSkipped = true
	}), nil
}
