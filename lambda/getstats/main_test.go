package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TODO add test case with 0 rows
// TODO test a range with dates far in the past
// TODO add test with gaps between days
// TODO update test cases so the dates are automatically updated

func TestRowsToBars(t *testing.T) {

	tt := []struct {
		name string
		rows []Row
		want []BarData
	}{

		{
			name: "empty set",
			rows: []Row{},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 0},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 0},
				{"2022-06-25 00:00:00 +0000 UTC", 0},
				{"2022-06-26 00:00:00 +0000 UTC", 0},
				{"2022-06-27 00:00:00 +0000 UTC", 0},
				{"2022-06-28 00:00:00 +0000 UTC", 0},
			},
		}, {
			name: "3 rows with 1 day",
			rows: []Row{
				{"2022-06-27T07:16:11.125Z", 1},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
			},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 0},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 0},
				{"2022-06-25 00:00:00 +0000 UTC", 0},
				{"2022-06-26 00:00:00 +0000 UTC", 0},
				{"2022-06-27 00:00:00 +0000 UTC", 6},
				{"2022-06-28 00:00:00 +0000 UTC", 0},
			},
		}, {
			name: "temp",
			rows: []Row{
				{"2022-06-24T07:16:11.125Z", 3},
				{"2022-06-25T09:19:55.747Z", 2},
				{"2022-06-26T21:44:54.892Z", 1},
				{"2022-06-28T09:44:54.892Z", 1},
			},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 0},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 3},
				{"2022-06-25 00:00:00 +0000 UTC", 2},
				{"2022-06-26 00:00:00 +0000 UTC", 1},
				{"2022-06-27 00:00:00 +0000 UTC", 0},
				{"2022-06-28 00:00:00 +0000 UTC", 1},
			},
		}, {
			name: "evenly distributed",
			rows: []Row{
				{"2022-06-24T07:16:11.125Z", 2},
				{"2022-06-25T09:19:55.747Z", 2},
				{"2022-06-26T09:19:55.747Z", 2},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-28T21:44:54.892Z", 2},
			},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 0},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 2},
				{"2022-06-25 00:00:00 +0000 UTC", 2},
				{"2022-06-26 00:00:00 +0000 UTC", 2},
				{"2022-06-27 00:00:00 +0000 UTC", 2},
				{"2022-06-28 00:00:00 +0000 UTC", 2},
			},
		}, {
			name: "gaps between days",
			rows: []Row{
				{"2022-06-22T07:16:11.125Z", 1},
				{"2022-06-25T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
			},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 1},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 0},
				{"2022-06-25 00:00:00 +0000 UTC", 2},
				{"2022-06-26 00:00:00 +0000 UTC", 0},
				{"2022-06-27 00:00:00 +0000 UTC", 3},
				{"2022-06-28 00:00:00 +0000 UTC", 0},
			},
		}, {
			name: "long period before",
			rows: []Row{
				{"2021-06-22T07:16:11.125Z", 1},
				{"2021-07-22T07:16:11.125Z", 1},
				{"2022-06-25T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
			},
			want: []BarData{
				{"2022-06-22 00:00:00 +0000 UTC", 0},
				{"2022-06-23 00:00:00 +0000 UTC", 0},
				{"2022-06-24 00:00:00 +0000 UTC", 0},
				{"2022-06-25 00:00:00 +0000 UTC", 2},
				{"2022-06-26 00:00:00 +0000 UTC", 0},
				{"2022-06-27 00:00:00 +0000 UTC", 3},
				{"2022-06-28 00:00:00 +0000 UTC", 0},
			},
		},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bars, err := GetBarData(tc.rows, 7)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(bars, tc.want) {
				fmt.Println("got:")
				for _, bar := range bars {
					fmt.Println(bar.Date, bar.Minutes)
				}
				fmt.Println("want:")
				for _, bar := range tc.want {
					fmt.Println(bar.Date, bar.Minutes)
				}

				t.Error("got != want")
			}
		})
	}

}

func TestGetWeeklyData(t *testing.T) {
	tt := []struct {
		name string
		rows []Row
		want []WeeklyData
	}{

		{
			name: "only monday",
			rows: []Row{
				{"2022-06-27T07:16:11.125Z", 1},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
			},
			want: []WeeklyData{
				{"Sun", 0},
				{"Mon", 6},
				{"Tue", 0},
				{"Wed", 0},
				{"Thu", 0},
				{"Fri", 0},
				{"Sat", 0},
			},
		}, {
			name: "only monday",
			rows: []Row{
				{"2022-06-26T07:16:11.125Z", 1},
				{"2022-06-27T07:16:11.125Z", 1},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
				{"2022-06-28T21:44:54.892Z", 3},
			},
			want: []WeeklyData{
				{"Sun", 1},
				{"Mon", 6},
				{"Tue", 3},
				{"Wed", 0},
				{"Thu", 0},
				{"Fri", 0},
				{"Sat", 0},
			},
		},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bars, err := GetWeeklyData(tc.rows)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(bars, tc.want) {
				fmt.Println("got:")
				for _, bar := range bars {
					fmt.Println(bar.Weekday, bar.Minutes)
				}
				fmt.Println("want:")
				for _, bar := range tc.want {
					fmt.Println(bar.Weekday, bar.Minutes)
				}

				t.Error("got != want")
			}
		})
	}

}

func TestTotals(t *testing.T) {
	tt := []struct {
		name string
		n    int
		rows []Row
		want int
	}{

		{
			name: "only monday",
			n:    7,
			rows: []Row{
				{"2021-06-27T07:16:11.125Z", 1},
				{"2022-06-27T07:16:11.125Z", 1},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
			},
			want: 6,
		}, {
			name: "only monday",
			n:    1,
			rows: []Row{
				{"2022-06-26T07:16:11.125Z", 1},
				{"2022-06-27T07:16:11.125Z", 1},
				{"2022-06-27T09:19:55.747Z", 2},
				{"2022-06-27T21:44:54.892Z", 3},
				{"2022-06-28T21:44:54.892Z", 3},
			},
			want: 3,
		},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			total, err := Totals(tc.rows, tc.n)
			if err != nil {
				t.Error(err)
			}

			if total != tc.want {
				fmt.Println("got:")
				t.Errorf("got: %d\twant: %d\n", total, tc.want)
			}
		})
	}
}
