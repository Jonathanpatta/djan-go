package djan_go

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

func GetLambdaHandler(c *Config) (func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error), error) {
	handler := func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		// If no name is provided in the HTTP request body, throw an error
		fmt.Println("req:", req.RawPath)
		adapter := gorillamux.NewV2(c.Router)
		resp, err := adapter.ProxyWithContext(ctx, req)

		return resp, err

	}

	return handler, nil
}
