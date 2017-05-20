package calendar

import (
	"errors"
	"time"
)

// -----------------------------------------------------------------------------

var (
	// ErrLocaleNotFound is raised when the locale is not supported
	ErrLocaleNotFound = errors.New("calendar: given locale not found")
)

// -----------------------------------------------------------------------------
type defaultCalendar struct {
	name   string
	year   int
	wdFunc WorkDayFunc
}

// New calendar built using given locale
func New(name string, year int) (Calendar, error) {
	// Check supported locale
	if _, ok := locales[name]; !ok {
		return nil, ErrLocaleNotFound
	}

	return &defaultCalendar{
		name:   name,
		year:   year,
		wdFunc: locales[name](year),
	}, nil
}

// -----------------------------------------------------------------------------
func (d *defaultCalendar) IsWorkingDay(day time.Time) (bool, bool, string) {
	return d.wdFunc(day)
}
