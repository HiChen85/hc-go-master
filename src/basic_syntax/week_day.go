package basic_syntax

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
	if SATURDAY <= w && w <= SATURDAY {
		return weekDayNames[w]
	}
	return "invalid weekday value"
}
