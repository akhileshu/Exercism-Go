package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// Schedule("7/25/2019 13:45:00")
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
// HasPassed("July 25, 2019 13:45:00")
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
// Description("7/25/2019 13:45:00")
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
// Description("7/25/2019 13:45:00")
// => 2020-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	// t, _ := time.Parse("2006-01-02 15:04:00", fmt.Sprintf("%v-09-15 00:00:00", time.Now().Year()))
	// return t
    // above approach worked , but below is simpler solution
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)

}
