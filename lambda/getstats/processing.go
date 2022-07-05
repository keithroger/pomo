package main

import (
	"fmt"
	"strconv"
	"time"
)

// Stats struct contains values for simple statistics
type Stats struct {
	Today, Week, Month, Year, All                string
	TodayAvg, WeekAvg, MonthAvg, YearAvg, AllAvg string
}

func newStats(rows []Row) Stats {
	all := 0

	for _, row := range rows {
		all += row.minutes
	}

	allAvg := float64(all) / float64(len(rows))

	today, todayAvg := totalAvgMin(rows, 1)
	week, weekAvg := totalAvgMin(rows, 7)
	month, monthAvg := totalAvgMin(rows, 30)
	year, yearAvg := totalAvgMin(rows, 365)

	return Stats{
		Today:    today,
		TodayAvg: todayAvg,
		Week:     week,
		WeekAvg:  weekAvg,
		Month:    month,
		MonthAvg: monthAvg,
		Year:     year,
		YearAvg:  yearAvg,
		All:      readableStr(all),
		AllAvg:   strconv.FormatFloat(allAvg, 'f', 2, 64),
	}
}

// GetBarData calculates the number of minutes studied during n days.
func getPeriodData(rows []Row, n int) []BarDatum {
	startDate := today().AddDate(0, 0, -n+1)
	startIdx := firstIdx(rows, startDate)

	// Trim out rows not in period
	if startIdx == -1 {
		rows = []Row{}
	} else {
		rows = rows[startIdx:]
	}

	// initialize dates
	data := make([]BarDatum, n)
	for i := range data {
		data[i].X = startDate.AddDate(0, 0, i).String()
	}

	// Counts number of minutes for each date
	minuteCount := make(map[string]int)
	for _, row := range rows {
		minuteCount[row.date.String()] += row.minutes
	}

	// Add counts from map counter and make dates into a readable format
	for i, d := range data {
		data[i].Y = minuteCount[data[i].X]
		data[i].X = d.X[5:7] + "/" + d.X[8:10]
	}

	return data
}

// getWeekDayData returns minute coutes by weekday for the last 30 days.
func getWeekDayData(rows []Row) []BarDatum {
	startDate := today().AddDate(0, 0, -30)
	startIdx := firstIdx(rows, startDate)

	// trim out dates not within 30 days
	if startIdx == -1 {
		rows = []Row{}
	} else {
		rows = rows[startIdx:]
	}

	// initialize data
	data := make([]BarDatum, 7)
	for i := range data {
		data[i].X = time.Weekday(i).String()
	}

	// count minutes in each weekday
	counts := make(map[string]int)
	for _, row := range rows {
		counts[row.date.Weekday().String()] += row.minutes
	}

	// add data from counts map
	for i := range data {
		data[i].Y = counts[data[i].X]
		// shorten weekday name
		data[i].X = data[i].X[:3]
	}

	return data
}

// totals gets the total minutes studied in the last n days
func totalAvgMin(rows []Row, n int) (string, string) {
	total := 0

	startDay := today().AddDate(0, 0, -n+1)
	startIdx := firstIdx(rows, startDay)
	if startIdx == -1 {
		return "0", "0"
	}

	// Only start counting from starting date
	rows = rows[startIdx:]

	for _, row := range rows {
		total += row.minutes
	}

	avg := float64(total) / float64(len(rows))

	return readableStr(total), strconv.FormatFloat(avg, 'f', 2, 64)
}

// readableStr takes number of minutes and converts it into a readable
// format (ex. 1d 12h 30m)
func readableStr(inMinutes int) string {
	days := inMinutes / (24 * 60)
	hours := inMinutes % (24 * 60) / 60
	minutes := inMinutes % 60

	if days > 0 {
		return fmt.Sprintf("%dd%dh%dm", days, hours, minutes)
	}

	if hours > 0 {
		return fmt.Sprintf("%dh%dm", hours, minutes)
	}

	return fmt.Sprintf("%dm", minutes)
}

func firstIdx(rows []Row, startDate time.Time) int {
	// check if the slice contains any value greater than date
	if len(rows) == 0 || startDate.After(rows[len(rows)-1].date) {
		return -1
	}

	lo, hi := 0, len(rows)-1
	for lo != hi {
		mid := (hi + lo) / 2
		if startDate.After(rows[mid].date) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

// dayOnly takes a date and truncates the measures smaller than the day.
func dayOnly(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

// Parse a ISOstring to time.Time without units smaller than days
func parseToDay(timeStr string) (time.Time, error) {
	t, err := time.Parse(isoLayout, timeStr)
	return dayOnly(t), err
}

func today() time.Time {
	return dayOnly(time.Now())
}
