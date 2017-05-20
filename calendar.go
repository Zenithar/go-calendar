package calendar

import "time"

// Calendar respresents default calendar contract
type Calendar interface {
	IsWorkingDay(time.Time) (bool, bool, string)
}
