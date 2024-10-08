package api

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func HandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("implement me")
}
