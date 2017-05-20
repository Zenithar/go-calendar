package calendar

import "time"

func frFR(year int) WorkDayFunc {
	// Calculate variable dates
	easterDay := easterDayMeus(year)
	easterMonDay := easterDay.Add(24 * time.Hour)
	ascensionDay := easterDay.Add(39 * 24 * time.Hour)
	pentecostDay := ascensionDay.Add(10 * 24 * time.Hour)
	pentecostMonDay := pentecostDay.Add(24 * time.Hour)

	// Return workingDay function for calendar
	return func(day time.Time) (bool, bool, string) {
		weekendDay := false

		// Filter by week days
		switch day.Weekday() {
		case time.Saturday, time.Sunday:
			weekendDay = true
		}

		// Filter fixed days by month
		switch day.Month() {
		case time.January:
			switch day.Day() {
			case 1:
				return false, weekendDay, "Jour de l'an"
			}
			//		case time.February:
			//		case time.March:
			//		case time.April:
		case time.May:
			switch day.Day() {
			case 1:
				return false, weekendDay, "Fête du travail"
			case 8:
				return false, weekendDay, "Victoire des alliés 1945"
			}
			//		case time.June:
		case time.July:
			switch day.Day() {
			case 14:
				return false, weekendDay, "Fête nationale"
			}
		case time.August:
			switch day.Day() {
			case 15:
				return false, weekendDay, "Assomption"
			}
			//		case time.September:
			//		case time.October:
		case time.November:
			switch day.Day() {
			case 1:
				return false, weekendDay, "Toussaint"
			case 11:
				return false, weekendDay, "Armistice 1918"
			}
		case time.December:
			switch day.Day() {
			case 25:
				return false, weekendDay, "Noël"
			}
		}

		// Filter calculated days
		if day.Month() == easterDay.Month() && day.Day() == easterDay.Day() {
			return false, weekendDay, "Pâques"
		}

		if day.Month() == easterMonDay.Month() && day.Day() == easterMonDay.Day() {
			return false, weekendDay, "Lundi de Pâques"
		}

		if day.Month() == ascensionDay.Month() && day.Day() == ascensionDay.Day() {
			return false, weekendDay, "Ascension"
		}

		if day.Month() == pentecostDay.Month() && day.Day() == pentecostDay.Day() {
			return false, weekendDay, "Pentecôte"
		}

		if day.Month() == pentecostMonDay.Month() && day.Day() == pentecostMonDay.Day() {
			return false, weekendDay, "Lundi de Pentecôte"
		}

		if weekendDay {
			return false, true, "Jour de WeekEnd"
		}

		return true, false, ""
	}
}

func init() {
	registerLocale("fr_FR", frFR)
}
