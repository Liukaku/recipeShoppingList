package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	services "github.com/liukaku/shoppingLambda/cmd/services"
	types "github.com/liukaku/shoppingLambda/cmd/utils"
)

func getRequestMethod(event types.RequestShape) string {
	return event.RequestContext.HTTP.Method
}

func handleRequest(ctx context.Context, event json.RawMessage) (string, error) {
	// TODO: use uuid for id go get github.com/google/uuid

	var jsonEvent types.RequestShape
	json.Unmarshal(event, &jsonEvent)

	method := getRequestMethod(jsonEvent)

	if method == "GET" {
		return services.GetRecipeById(jsonEvent)
	} else if method == "POST" {
		return services.CreateRecipe(jsonEvent)
	} else {
		fmt.Println("default")
		return "Method not supported", nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
