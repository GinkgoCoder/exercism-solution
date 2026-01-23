package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func FindAllWeekdaysInMonth(wDay time.Weekday, monnth time.Month, year int) []int {
	firstDay := time.Date(year, monnth, 1, 0, 0, 0, 0, time.UTC)
	diff := (int(wDay-firstDay.Weekday()) + 7) % 7
	weekDay := firstDay.Add(time.Duration(diff) * 24 * time.Hour)
	allWeekDays := []int{}
	for weekDay.Month() == monnth {
		allWeekDays = append(allWeekDays, weekDay.Day())
		weekDay = weekDay.Add(7 * time.Hour * 24)
	}
	return allWeekDays
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	weekDays := FindAllWeekdaysInMonth(wDay, month, year)
	switch wSched {
	case First, Second, Third, Fourth:
		return weekDays[int(wSched)]
	case Last:
		return weekDays[len(weekDays)-1]
	case Teenth:
		for _, d := range weekDays {
			if d >= 13 && d <= 19 {
				return d
			}
		}
		return -1
	default:
		return -1
	}
}
