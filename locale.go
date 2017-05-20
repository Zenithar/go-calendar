package calendar

import "time"

// WorkDayFunc is the function contract that implements the localised calendar
type WorkDayFunc func(time.Time) (bool, bool, string)

// -----------------------------------------------------------------------------
var (
	locales map[string]func(int) WorkDayFunc
)

func registerLocale(name string, wd func(int) WorkDayFunc) {
	// Lazy initialization
	if locales == nil {
		locales = map[string]func(int) WorkDayFunc{}
	}
	locales[name] = wd
}
