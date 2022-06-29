package main

import (
	"context"
	"fmt"
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
	inLayout  = "2006-01-02T15:04:05.000Z"
	outLayout = "2006/01/02"
)

type Event struct {
	Username string `json:"username"`
}

type BarData struct {
	Date    string `json:"date"` // save to unix time
	Minutes int    `json:"minutes"`
}

type Row struct {
	Timestamp string `json:"timestamp"`
	Minutes   int    `json:"minutes"`
}

type Response struct {
	bar []BarData
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

func parse(timeStr string) (time.Time, error) {
	return time.Parse(inLayout, timeStr)

}

func HandleRequest(ctx context.Context, username Event) (string, error) {

	out, err := ddb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("pomo-study-user-stats"),
		KeyConditionExpression: aws.String("username = :username"),
		ConsistentRead:         aws.Bool(true),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":username": &types.AttributeValueMemberS{Value: username.Username},
		},
	})

	if err != nil {
		panic(err)
	}

	var rows []Row
	err = attributevalue.UnmarshalListOfMaps(out.Items, &rows)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Hello %s!", username.Username), nil
}

// TODO user userinput date otherwise it will be according to servers date
// Returns midnight representing the begining of the day.
// Values can be used as discreete dates.
// Ignores time zone.
func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func getBarData(rows []Row, n int) ([]BarData, error) {
	// get todays date at midnight
	today := midnight(time.Now())

	// find the first index
	// TODO optimize with binary search
	for i, row := range rows {

		rowDate, err := parse(row.Timestamp)
		if err != nil {
			return []BarData{}, err
		}
		rowDate = midnight(rowDate)

		firstDay := today.AddDate(0, 0, -n+1)
		if rowDate.Add(time.Second).After(firstDay) {
			rows = rows[i:]
			break
		}
	}

	// testing
	fmt.Println("Rows left")
	for _, row := range rows {
		fmt.Println(row)
	}

	// initialize dates
	bars := make([]BarData, n)
	for i := range bars {
		date := today.AddDate(0, 0, -n+i+1).String()
		bars[i].Date = date
	}

	// Add minutes to bars
	minuteCount := make(map[string]int)
	for _, row := range rows {

		date, err := parse(row.Timestamp)
		if err != nil {
			return []BarData{}, err
		}
		date = midnight(date)

		minuteCount[date.String()] += row.Minutes

		// i=0 : -7+0+2=-5
		// i=1 : -7+1+2=-4
		// i=2 : -7+2+2=-3
		// i=3 : -7+3+2=-2
		// i=4 : -7+4+2=-1
		// i=5 : -7+5+2= 0
		// i=6 : -7+6+2=+1
		// fmt.Printf("idx: %d\tdate: %s\n", barIdx, today.AddDate(0, 0, -n+barIdx+2).String())

		// for date.After(today.AddDate(0, 0, -n+barIdx+2)) {
		// 	barIdx++
		// }

		// bars[barIdx].Minutes += row.Minutes
	}

	for i := range bars {
		bars[i].Minutes = minuteCount[bars[i].Date]
	}

	return bars, nil
}

func main() {
	lambda.Start(HandleRequest)
}
