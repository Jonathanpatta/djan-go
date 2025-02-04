package main

import (
	"fmt"
	djan_go "github.com/Jonathanpatta/djan-go"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	c, err := djan_go.NewDefaultConfig()
	if err != nil {
		fmt.Println(err)
	}

	handler, _ := djan_go.GetLambdaHandler(c)

	lambda.Start(handler)
}
