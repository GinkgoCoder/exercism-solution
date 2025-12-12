package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	if t, err := time.Parse("1/2/2006 15:04:05", date); err != nil {
        return time.Time{}
    } else {
        return t
    }
    
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    h, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return h.Hour() >= 12 && h.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2024, time.September, 15, 0,0,0,0, time.UTC).AddDate(1, 0, 0)
}
