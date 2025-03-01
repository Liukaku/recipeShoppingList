package services

import (
	"encoding/json"
	"fmt"

	"github.com/liukaku/shoppingLambda/cmd/aws"
	types "github.com/liukaku/shoppingLambda/cmd/utils"
)

func GetAllRecipes() (string, error) {

	got := aws.GetAll()
	fmt.Println(got)
	combined := combineIncredients(got)

	response, err := json.Marshal(combined)
	if err != nil {
		panic(err)
	}
	fmt.Println("response", string(response))

	return string(response), nil
}

func combineIncredients(ingredients []types.Recipe) []types.Ingredient {
	// combine ingredients
	// if the same ingredient is found, combine the amount
	// if the same ingredient is found with different units, convert to the first unit

	retVal := []types.Ingredient{}
	for _, v := range ingredients {
		for _, i := range v.Ingredients {
			found := false
			for j, r := range retVal {
				if r.Name == i.Name {
					retVal[j].Amount += i.Amount
					found = true
					break
				}
			}
			if !found {
				convertedAmount, unit := convertImperialToMetric(i.Amount, i.Unit)
				retVal = append(retVal,
					types.Ingredient{
						ID:     i.ID,
						Name:   i.Name,
						Amount: int(convertedAmount),
						Unit:   unit,
					},
				)
			}
		}
	}
	return retVal
}

func convertImperialToMetric(amount int, unit string) (float64, string) {
	if unit == "cup" {
		return float64(amount) * 240, "ml"
	} else if unit == "tbsp" {
		return float64(amount) * 15, "ml"
	} else if unit == "tsp" {
		return float64(amount) * 5, "ml"
	} else if unit == "oz" {
		return float64(amount) * 30, "ml"
	} else if unit == "lb" {
		return float64(amount) * 454, "g"
	} else if unit == "inch" {
		return float64(amount) * 2.54, "cm"
	} else if unit == "foot" {
		return float64(amount) * 30.48, "cm"
	} else if unit == "yard" {
		return float64(amount) * 91.44, "cm"
	} else if unit == "mile" {
		return float64(amount) * 1.609, "km"
	} else if unit == "gallon" {
		return float64(amount) * 3.785, "l"
	} else if unit == "pint" {
		return float64(amount) * 0.473, "l"
	} else if unit == "quart" {
		return float64(amount) * 0.946, "l"
	} else if unit == "fl oz" {
		return float64(amount) * 29.574, "ml"
	}
	return float64(amount), unit
}
