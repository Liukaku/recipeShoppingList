package services

import (
	"encoding/json"
	"fmt"

	"github.com/liukaku/shoppingLambda/cmd/aws"
	types "github.com/liukaku/shoppingLambda/cmd/utils"
)

func CreateRecipe(jsonEvent types.RequestShape) (string, error) {
	fmt.Println("POST")
	var recipe types.Recipe
	json.Unmarshal([]byte(jsonEvent.Body), &recipe)
	fmt.Println("recipe", recipe)

	err := aws.Upsert(recipe)

	if err != nil {
		panic(err)
	}

	return "Success", nil
}
