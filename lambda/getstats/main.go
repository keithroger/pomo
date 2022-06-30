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

type WeeklyData struct {
	Weekday string `json:"weekday"`
	Minutes int    `json:"minutes"`
}

type Row struct {
	Timestamp string `json:"timestamp"`
	Minutes   int    `json:"minutes"`
}

// TODO make numbers into formated strings
type Response struct {
	Bar7Day                       []BarData `json:"7 Days"`
	Bar14Day                      []BarData `json:"14 Days"`
	Bar30Day                      []BarData `json:"30 Days"`
	WeekdayData                   []WeeklyData
	Today, Week, Month, Year, All int
	AllAvg                        float32
}

var (
	ddb *dynamodb.Client
)

// TODO make a struct to convert rows to have time.Time objects
// TODO clean up code and efficiency
// TODO add averages for different intervals

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

func HandleRequest(ctx context.Context, username Event) (Response, error) {

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

	// TODO refactor code to prevent needing to handle errors every time
	Bar7Day, err := GetBarData(rows, 7)
	checkErr(err)
	Bar14Day, err := GetBarData(rows, 14)
	checkErr(err)
	Bar30Day, err := GetBarData(rows, 30)
	checkErr(err)
	WeekdayData, err := GetWeeklyData(rows)
	checkErr(err)
	// TODO put all the totals together in a struct
	Today, err := TotalMinutes(rows, 1)
	checkErr(err)
	Week, err := TotalMinutes(rows, 7)
	checkErr(err)
	Month, err := TotalMinutes(rows, 30)
	checkErr(err)
	// TODO account for leap years
	Year, err := TotalMinutes(rows, 365)
	checkErr(err)
	All, err := TotalMinutes(rows, len(rows))
	checkErr(err)

	resp := Response{
		Bar7Day:     Bar7Day,
		Bar14Day:    Bar14Day,
		Bar30Day:    Bar30Day,
		WeekdayData: WeekdayData,
		Today:       Today,
		Week:        Week,
		Month:       Month,
		Year:        Year,
		All:         All,
		AllAvg:      float32(All) / float32(len(rows)),
	}

	return resp, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO user userinput date otherwise it will be according to servers date
// Returns midnight representing the begining of the day.
// Values can be used as discreete dates.
// Ignores time zone.
func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// GetBarData calculates the number of minutes studied during n days.
func GetBarData(rows []Row, n int) ([]BarData, error) {
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

	}

	for i := range bars {
		bars[i].Minutes = minuteCount[bars[i].Date]
	}

	return bars, nil
}

func GetWeeklyData(rows []Row) ([]WeeklyData, error) {
	bars, err := GetBarData(rows, 30)
	if err != nil {
		return []WeeklyData{}, err
	}

	// initialize days of the week
	weeklyBars := make([]WeeklyData, 7)
	for i := 0; i < 7; i++ {
		weeklyBars[i].Weekday = time.Weekday(i).String()[:3]
	}

	// add minutes to the corisponding weekday count
	for _, bar := range bars {
		date, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", bar.Date)
		if err != nil {
			return []WeeklyData{}, err
		}
		weekday := date.Weekday()
		weeklyBars[int(weekday)].Minutes += bar.Minutes
	}

	return weeklyBars, nil
}

// totals gets the total minutes studied in the last n days
func TotalMinutes(rows []Row, n int) (int, error) {
	total := 0
	threshold := midnight(time.Now()).AddDate(0, 0, -n+1)

	for i := len(rows) - 1; i >= 0; i-- {
		date, err := parse(rows[i].Timestamp)
		if err != nil {
			return 0, err
		}
		if date.After(threshold) {
			total += rows[i].Minutes
		} else {
			break
		}

	}

	return total, nil
}

func main() {
	lambda.Start(HandleRequest)
}
