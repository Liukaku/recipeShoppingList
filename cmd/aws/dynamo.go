package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	listTypes "github.com/liukaku/shoppingLambda/cmd/utils"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func loadDb() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))

	// return error and handle it as a return value
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}

func Get(id string) listTypes.Recipe {
	svc := loadDb()
	idMap := map[string]string{"ID": id}
	key, err := attributevalue.MarshalMap(idMap)
	if err != nil {
		panic("Unable to marshal key")
	}
	data, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("shopping-recipe"),
		Key:       key,
	},
	)
	if err != nil {
		panic("Unable to get item")
	}

	item := listTypes.Recipe{}

	attributevalue.UnmarshalMap(data.Item, &item)

	return item
}

func GetAll() []listTypes.Recipe {
	svc := loadDb()
	data, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("shopping-recipe"),
	})

	if err != nil {
		panic("Unable to scan table")
	}

	items := []listTypes.Recipe{}

	for _, i := range data.Items {
		item := listTypes.Recipe{}
		attributevalue.UnmarshalMap(i, &item)
		items = append(items, item)
	}

	return items
}

func Upsert(recipe listTypes.Recipe) error {
	svc := loadDb()
	fmt.Println("Redicpe ID: ", recipe.ID)
	fmt.Println("Recipe Ingredients Id: ", recipe.Ingredients[0].ID)
	data, err := attributevalue.MarshalMap(recipe)

	// return error and handle it as a return value
	if err != nil {
		panic("Unable to marshal recipe")
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("shopping-recipe"),
		Item:      data,
	})

	if err != nil {
		fmt.Println(err)
		panic("Unable to put item")
	}

	return nil
}
