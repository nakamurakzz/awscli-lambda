package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: "Hello, " + event.Name}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
