package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConvertRows(t *testing.T) {
	now := time.Now()
	ascendingDates := []time.Time{
		now.AddDate(0, 0, -3),
		now.AddDate(0, 0, -2),
		now.AddDate(0, 0, -1),
	}

	// test table
	tt := []struct {
		name    string
		inRows  []QueryRow
		outRows []Row
	}{
		{
			name: "Constant value",
			inRows: []QueryRow{
				{now.Format(isoLayout), 1},
				{now.Format(isoLayout), 1},
				{now.Format(isoLayout), 1},
			},
			outRows: []Row{
				{today(), 1},
				{today(), 1},
				{today(), 1},
			},
		}, {
			name: "Ascending",
			inRows: []QueryRow{
				{ascendingDates[0].Format(isoLayout), 1},
				{ascendingDates[1].Format(isoLayout), 1},
				{ascendingDates[2].Format(isoLayout), 1},
			},
			outRows: []Row{
				{dayOnly(ascendingDates[0]), 1},
				{dayOnly(ascendingDates[1]), 1},
				{dayOnly(ascendingDates[2]), 1},
			},
		}, {
			name: "Different minutes",
			inRows: []QueryRow{
				{now.Format(isoLayout), 1},
				{now.Format(isoLayout), 2},
				{now.Format(isoLayout), 3},
			},
			outRows: []Row{
				{today(), 1},
				{today(), 2},
				{today(), 3},
			},
		},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, err := newRows(tc.inRows)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(data, tc.outRows) {

				// Create error string
				errStr := "got:\n"
				for _, datum := range data {
					errStr += fmt.Sprintln(datum.date, datum.minutes)
				}
				errStr += "want:\n"
				for _, datum := range tc.outRows {
					errStr += fmt.Sprintln(datum.date, datum.minutes)
				}

				t.Error("got != want\n" + errStr)
			}
		})
	}

}

func TestGetPeriodData(t *testing.T) {
	dates := []string{
		today().AddDate(0, 0, -6).String(),
		today().AddDate(0, 0, -5).String(),
		today().AddDate(0, 0, -4).String(),
		today().AddDate(0, 0, -3).String(),
		today().AddDate(0, 0, -2).String(),
		today().AddDate(0, 0, -1).String(),
		today().AddDate(0, 0, 0).String(),
	}

	for i, d := range dates {
		dates[i] = d[5:7] + "/" + d[8:10]
	}

	tt := []struct {
		name string
		rows []Row
		want []BarDatum
	}{

		{
			name: "Empty set",
			rows: []Row{},
			want: []BarDatum{
				{dates[0], 0},
				{dates[1], 0},
				{dates[2], 0},
				{dates[3], 0},
				{dates[4], 0},
				{dates[5], 0},
				{dates[6], 0},
			},
		}, {
			name: "Constant Date",
			rows: []Row{
				{today(), 1},
				{today(), 1},
				{today(), 1},
			},
			want: []BarDatum{
				{dates[0], 0},
				{dates[1], 0},
				{dates[2], 0},
				{dates[3], 0},
				{dates[4], 0},
				{dates[5], 0},
				{dates[6], 3},
			},
		}, {
			name: "One Date",
			rows: []Row{
				{today().AddDate(0, 0, -5), 2},
			},
			want: []BarDatum{
				{dates[0], 0},
				{dates[1], 2},
				{dates[2], 0},
				{dates[3], 0},
				{dates[4], 0},
				{dates[5], 0},
				{dates[6], 0},
			},
		}, {
			name: "Date Before Period",
			rows: []Row{
				{today().AddDate(0, -1, 0), 2},
			},
			want: []BarDatum{
				{dates[0], 0},
				{dates[1], 0},
				{dates[2], 0},
				{dates[3], 0},
				{dates[4], 0},
				{dates[5], 0},
				{dates[6], 0},
			},
		}, {
			name: "Gaps Between Dates",
			rows: []Row{
				{today().AddDate(0, 0, 0), 2},
				{today().AddDate(0, 0, -2), 2},
				{today().AddDate(0, 0, -4), 2},
			},
			want: []BarDatum{
				{dates[0], 0},
				{dates[1], 0},
				{dates[2], 2},
				{dates[3], 0},
				{dates[4], 2},
				{dates[5], 0},
				{dates[6], 2},
			},
		},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data := getPeriodData(tc.rows, 7)

			if !reflect.DeepEqual(data, tc.want) {

				// Create error string
				errStr := "got:\n"
				for _, datum := range data {
					errStr += fmt.Sprintln(datum.X, datum.Y)
				}
				errStr += "want:\n"
				for _, datum := range tc.want {
					errStr += fmt.Sprintln(datum.X, datum.Y)
				}

				t.Error("got != want\n" + errStr)
			}
		})
	}

}

func TestGetWeekDayData(t *testing.T) {
	tt := []struct {
		name string
		rows []Row
		want []BarDatum
	}{

		{
			name: "Empty Set",
			rows: []Row{},
			want: []BarDatum{
				{"Sun", 0},
				{"Mon", 0},
				{"Tue", 0},
				{"Wed", 0},
				{"Thu", 0},
				{"Fri", 0},
				{"Sat", 0},
			},
		}, {
			name: "Constant Value",
			rows: []Row{
				{today().AddDate(0, 0, 0), 2},
				{today().AddDate(0, 0, -1), 2},
				{today().AddDate(0, 0, -2), 2},
				{today().AddDate(0, 0, -3), 2},
				{today().AddDate(0, 0, -4), 2},
				{today().AddDate(0, 0, -5), 2},
				{today().AddDate(0, 0, -6), 2},
			},
			want: []BarDatum{
				{"Sun", 2},
				{"Mon", 2},
				{"Tue", 2},
				{"Wed", 2},
				{"Thu", 2},
				{"Fri", 2},
				{"Sat", 2},
			},
		}, {
			name: "Dates Older than 30 Days",
			rows: []Row{
				{today().AddDate(0, -2, 0), 2},
				{today().AddDate(-1, 0, -1), 2},
				{today().AddDate(0, -3, -2), 2},
			},
			want: []BarDatum{
				{"Sun", 0},
				{"Mon", 0},
				{"Tue", 0},
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

			data := getWeekDayData(tc.rows)

			if !reflect.DeepEqual(data, tc.want) {
				// Create error string
				errStr := "got:\n"
				for _, datum := range data {
					errStr += fmt.Sprintln(datum.X, datum.Y)
				}
				errStr += "want:\n"
				for _, datum := range tc.want {
					errStr += fmt.Sprintln(datum.X, datum.Y)
				}

				t.Error("got != want\n" + errStr)
			}
		})
	}

}

func TestNewStats(t *testing.T) {

	inputRows := []Row{
		{today().AddDate(-2, 0, 0), 4},
		{today().AddDate(-1, 0, -3), 4},
		{today().AddDate(0, -1, -1), 4},
		{today().AddDate(0, -1, 0), 4},
		{today().AddDate(0, 0, -20), 12},
		{today().AddDate(0, 0, -2), 2},
		{today().AddDate(0, 0, -2), 2},
		{today().AddDate(0, 0, -1), 2},
		{today().AddDate(0, 0, 0), 2},
	}

	want := Stats{
		Today:    "2m",
		TodayAvg: "2.00",
		Week:     "8m",
		WeekAvg:  "2.00",
		Month:    "20m",
		MonthAvg: "4.00",
		Year:     "28m",
		YearAvg:  "4.00",
		All:      "36m",
		AllAvg:   "4.00",
	}

	stats := newStats(inputRows)

	if !reflect.DeepEqual(stats, want) {
		// Create error string
		errStr := "got:\n"
		errStr += fmt.Sprintf("%v\n", stats)
		errStr += "want:\n"
		errStr += fmt.Sprintf("%v\n", want)

		t.Error("got != want\n" + errStr)

	}
}

func TestReadableStr(t *testing.T) {
	tt := []struct {
		name    string
		minutes int
		want    string
	}{
		{"One Minutes", 1, "1m"},
		{"Two Minutes", 2, "2m"},
		{"One Hour One Minute", 61, "1h1m"},
		{"One Day One Minute", 24*60 + 1, "1d0h1m"},
	}

	for _, tc := range tt {
		tc := tc // capture variable

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			str := readableStr(tc.minutes)

			if tc.want != str {
				t.Errorf("got: %s \t want: %s\n", str, tc.want)
			}
		})
	}
}
