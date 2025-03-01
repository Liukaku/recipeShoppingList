package services

import (
	"encoding/json"
	"fmt"

	"github.com/liukaku/shoppingLambda/cmd/aws"
)

func GetRecipeById(id string) (string, error) {
	fmt.Println("GET")
	got := aws.Get(id)
	fmt.Println(got)
	response, err := json.Marshal(got)
	if err != nil {
		panic(err)
	}

	fmt.Println("response", string(response))

	return string(response), nil
}
