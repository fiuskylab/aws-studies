package main

import (
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/fiuskylab/aws-studies/lambda/handler"
)

func main() {
	runtime.Start(handler.HandleRequest)
}
