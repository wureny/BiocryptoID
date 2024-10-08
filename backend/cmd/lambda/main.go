package main

import (
	"biocryptoID/Internal/api"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.HandlerRequest)
}
