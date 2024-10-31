package awsconfig

import (
	"biocryptoID/flags"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewAWSConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			flags.AWSAccessKeyID,
			flags.AWSSecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
