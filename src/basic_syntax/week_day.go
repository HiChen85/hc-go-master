package basic_syntax

import "fmt"

type Weekday int

const (
	SUNDAY Weekday = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var weekDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func (w Weekday) String() string {
	if SUNDAY <= w && w <= SATURDAY {
		return weekDayNames[w]
	}
	return fmt.Sprintf("Invalid Weekday(%d)", w)
}
