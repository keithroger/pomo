package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Event struct {
	Username string `json:"username"`
}

type Row struct {
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}

var (
	ddb *dynamodb.Client
)

// Setup connection to dynamodb
func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-west-1"
		return nil
	})

	if err != nil {
		panic(err)
	}

	ddb = dynamodb.NewFromConfig(cfg)
}

func HandleRequest(ctx context.Context, username Event) error {

	// Search for items in stats table
	statsDDB, err := ddb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("pomo-study-user-stats"),
		KeyConditionExpression: aws.String("username = :username"),
		ConsistentRead:         aws.Bool(true),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":username": &types.AttributeValueMemberS{Value: username.Username},
		},
	})
	if err != nil {
		return err
	}

	// Unmarshall rows
	var rows []Row
	err = attributevalue.UnmarshalListOfMaps(statsDDB.Items, &rows)
	if err != nil {
		return err
	}

	// Delete items from stats table
	for _, row := range rows {
		_, err := ddb.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String("pomo-study-user-stats"),
			Key: map[string]types.AttributeValue{
				"username":  &types.AttributeValueMemberS{Value: row.Username},
				"timestamp": &types.AttributeValueMemberN{Value: row.Timestamp},
			},
		})
		if err != nil {
			return err
		}
	}

	// Delete entry for settings table
	_, err = ddb.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("pomo-study-user-settings"),
		Key: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: username.Username},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
