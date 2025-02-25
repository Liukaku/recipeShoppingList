package services

import (
	"encoding/json"
	"fmt"

	"github.com/liukaku/shoppingLambda/cmd/aws"
	types "github.com/liukaku/shoppingLambda/cmd/utils"
)

func GetRecipeById(jsonEvent types.RequestShape) (string, error) {
	fmt.Println("GET")
	id := jsonEvent.RawPath
	// remove initial slash from the endpoint, otherwise it will be part of the id
	id = id[1:]
	fmt.Println("id", id)
	got := aws.Get(id)
	fmt.Println(got)
	response, err := json.Marshal(got)
	if err != nil {
		panic(err)
	}

	fmt.Println("response", string(response))

	return string(response), nil
}
