package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "github.com/aws/aws-lambda-go/lambda"
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
		id := jsonEvent.RawPath
		// do we to have a separate route for simplifying the result for a shopping list
		// or should it return both a simplified and everything list?
		if id == "/all" {
			return services.GetAllRecipes()
		} else {
			return services.GetRecipeById(id)
		}
	} else if method == "POST" {
		return services.CreateRecipe(jsonEvent)
	} else {
		fmt.Println("default")
		return "Method not supported", nil
	}
}

func main() {
	// event := `{ "requestContext": { "http": { "method": "GET" } }, "rawPath": "/all" }`
	// handleRequest(context.Background(), []byte(event))
	lambda.Start(handleRequest)
}
