package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TableName = "pomo-study-user-stats"
	isoLayout = "2006-01-02T15:04:05.999Z"
)

type Event struct {
	Username string `json:"username"`
}

type QueryRow struct {
	Timestamp string `json:"timestamp"`
	Minutes   int    `json:"minutes"`
}

type Row struct {
	date    time.Time
	minutes int
}

func newRows(inRows []QueryRow) ([]Row, error) {
	outRows := make([]Row, len(inRows))

	for i, row := range inRows {
		date, err := parseToDay(row.Timestamp)
		if err != nil {
			return []Row{}, err
		}

		outRows[i] = Row{date, row.Minutes}
	}

	return outRows, nil
}

type BarDatum struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type Period struct {
	Days7  []BarDatum `json:"7 Days"`
	Days14 []BarDatum `json:"14 Days"`
	Days30 []BarDatum `json:"30 Days"`
}

// TODO look into returning a status code
type Response struct {
    PeriodData  Period `json:"period"`
    WeekDayData []BarDatum `json:"weekDayData"`
    Statistics  Stats `json:"stats"`
}

var (
	ddb *dynamodb.Client
)

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

func HandleRequest(ctx context.Context, username Event) Response {

	// Get rows from DynamoDB
	queryOut, err := ddb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(TableName),
		KeyConditionExpression: aws.String("username = :username"),
		ConsistentRead:         aws.Bool(true),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":username": &types.AttributeValueMemberS{Value: username.Username},
		},
	})
	if err != nil {
		panic(err)
	}

	// Unmarshall DynamoDB rows
	var qRows []QueryRow
	err = attributevalue.UnmarshalListOfMaps(queryOut.Items, &qRows)
	if err != nil {
		panic(err)
	}

	// If empty return
	if len(qRows) == 0 {
		return Response{}
	}

	// Convert rows to a more useable format
	rows, err := newRows(qRows)
	if err != nil {
		panic(err)
	}

	return Response{
		PeriodData: Period{
			Days7:  getPeriodData(rows, 7),
			Days14: getPeriodData(rows, 14),
			Days30: getPeriodData(rows, 30),
		},
		WeekDayData: getWeekDayData(rows),
		Statistics:  newStats(rows),
	}

}

func main() {
	lambda.Start(HandleRequest)
}
