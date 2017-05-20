package calendar_test

import (
	"testing"
	"time"

	"github.com/zenithar/calendar"
)

func TestInvalidLocale(t *testing.T) {
	_, err := calendar.New("toto", 2017)
	if err == nil {
		t.Error("Error should be raised on invalid locale usage.")
		t.FailNow()
	}
}

func TestFrenchCalendar(t *testing.T) {
	cal, err := calendar.New("fr_FR", 2017)
	if err != nil {
		t.Error("Unable to initialize calendar with 'fr_FR' locale.")
		t.FailNow()
	}

	expectedWorkingDayCount := 251
	expectedWeekendDayCount := 105
	expectedHolyDayCount := 13

	expectedHolydays := map[string]time.Time{
		"Jour de l'an":             time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
		"Pâques":                   time.Date(2017, time.April, 16, 0, 0, 0, 0, time.UTC),
		"Lundi de Pâques":          time.Date(2017, time.April, 17, 0, 0, 0, 0, time.UTC),
		"Fête du travail":          time.Date(2017, time.May, 1, 0, 0, 0, 0, time.UTC),
		"Victoire des alliés 1945": time.Date(2017, time.May, 8, 0, 0, 0, 0, time.UTC),
		"Ascension":                time.Date(2017, time.May, 25, 0, 0, 0, 0, time.UTC),
		"Pentecôte":                time.Date(2017, time.June, 4, 0, 0, 0, 0, time.UTC),
		"Lundi de Pentecôte":       time.Date(2017, time.June, 5, 0, 0, 0, 0, time.UTC),
		"Fête nationale":           time.Date(2017, time.July, 14, 0, 0, 0, 0, time.UTC),
		"Assomption":               time.Date(2017, time.August, 15, 0, 0, 0, 0, time.UTC),
		"Toussaint":                time.Date(2017, time.November, 1, 0, 0, 0, 0, time.UTC),
		"Armistice 1918":           time.Date(2017, time.November, 11, 0, 0, 0, 0, time.UTC),
		"Noël":                     time.Date(2017, time.December, 25, 0, 0, 0, 0, time.UTC),
	}

	// Iterate over each day of a year
	workedDayCount := 0
	weekendDayCount := 0
	holyDayCount := 0

	current := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 365; i++ {
		worked, weekendDay, name := cal.IsWorkingDay(current)
		if worked {
			workedDayCount++
		}
		if weekendDay {
			weekendDayCount++
		}

		if date, ok := expectedHolydays[name]; ok {
			if worked {
				t.Errorf("No work on '%s', because it's '%s'.", current, name)
				t.FailNow()
			}
			if date != current {
				t.Errorf("Invalid '%s' date, should be '%s' not '%s'.", name, date, current)
				t.FailNow()
			}

			holyDayCount++
		} else {
			if len(name) > 0 && name != "Jour de WeekEnd" {
				t.Errorf("Invalid holyday '%s' on %s", name, current)
			}
		}

		if current.Weekday() == time.Saturday {
			if worked {
				t.Errorf("No work on '%s', because it's Saturday.", current)
			}
		}

		if current.Weekday() == time.Sunday {
			if worked {
				t.Errorf("No work on '%s', because it's Sunday.", current)
			}
		}

		//fmt.Printf("%s, %v, %v, %s\n", current, worked, weekendDay, name)

		current = current.Add(24 * time.Hour)
	}

	if workedDayCount != expectedWorkingDayCount {
		t.Errorf("Worked day count is invalid '%d', expected '%d'.", workedDayCount, expectedWorkingDayCount)
	}
	if weekendDayCount != expectedWeekendDayCount {
		t.Errorf("Weekend day count is invalid '%d', expected '%d'.", weekendDayCount, expectedWeekendDayCount)
	}
	if holyDayCount != expectedHolyDayCount {
		t.Errorf("Holy day count is invalid '%d', expected '%d'.", holyDayCount, expectedHolyDayCount)
	}

}
