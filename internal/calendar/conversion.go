package calendar

import "fmt"

// CivilDate is a date in a named civil calendar. Year numbering is positive
// CE only; the conversion helpers do not model a jurisdiction's historical
// calendar-reform cutover.
type CivilDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// GregorianToJulian converts a proleptic Gregorian date to the corresponding
// proleptic Julian date. The returned day number is the integer Julian Day
// Number for the shared civil day; Julian dates begin at noon UT1.
func GregorianToJulian(year, month, day int) (CivilDate, int, error) {
	if !validCivilDate(year, month, day, false) {
		return CivilDate{}, 0, fmt.Errorf("invalid Gregorian date %04d-%02d-%02d", year, month, day)
	}
	jdn := gregorianJDN(year, month, day)
	converted := jdnToJulian(jdn)
	if converted.Year < 1 || converted.Year > 9999 {
		return CivilDate{}, 0, fmt.Errorf("converted Julian year is outside 1..9999")
	}
	return converted, jdn, nil
}

// JulianToGregorian converts a proleptic Julian date to the corresponding
// proleptic Gregorian date. The returned day number is the integer Julian Day
// Number for the shared civil day; Julian dates begin at noon UT1.
func JulianToGregorian(year, month, day int) (CivilDate, int, error) {
	if !validCivilDate(year, month, day, true) {
		return CivilDate{}, 0, fmt.Errorf("invalid Julian date %04d-%02d-%02d", year, month, day)
	}
	jdn := julianJDN(year, month, day)
	converted := jdnToGregorian(jdn)
	if converted.Year < 1 || converted.Year > 9999 {
		return CivilDate{}, 0, fmt.Errorf("converted Gregorian year is outside 1..9999")
	}
	return converted, jdn, nil
}

func validCivilDate(year, month, day int, julian bool) bool {
	if year < 1 || year > 9999 || month < 1 || month > 12 || day < 1 {
		return false
	}
	days := DaysInMonth(year, month)
	if julian {
		days = julianDaysInMonth(year, month)
	}
	return day <= days
}

func julianDaysInMonth(year, month int) int {
	switch month {
	case 2:
		if year%4 == 0 {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func gregorianJDN(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

func julianJDN(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - 32083
}

func jdnToGregorian(jdn int) CivilDate {
	a := jdn + 32044
	b := (4*a + 3) / 146097
	c := a - (146097*b)/4
	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153
	day := e - (153*m+2)/5 + 1
	month := m + 3 - 12*(m/10)
	year := 100*b + d - 4800 + m/10
	return CivilDate{Year: year, Month: month, Day: day}
}

func jdnToJulian(jdn int) CivilDate {
	c := jdn + 32082
	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153
	day := e - (153*m+2)/5 + 1
	month := m + 3 - 12*(m/10)
	year := d - 4800 + m/10
	return CivilDate{Year: year, Month: month, Day: day}
}
