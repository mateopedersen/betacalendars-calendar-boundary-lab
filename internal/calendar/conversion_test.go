package calendar

import "testing"

func TestGregorianJulianKnownCalendarDates(t *testing.T) {
	cases := []struct {
		gregorian CivilDate
		julian    CivilDate
	}{
		{CivilDate{2027, 1, 1}, CivilDate{2026, 12, 19}},
		{CivilDate{1582, 10, 15}, CivilDate{1582, 10, 5}},
		{CivilDate{1900, 3, 13}, CivilDate{1900, 2, 29}},
		{CivilDate{2000, 3, 13}, CivilDate{2000, 2, 29}},
	}
	for _, tc := range cases {
		gotJulian, gregorianJDN, err := GregorianToJulian(tc.gregorian.Year, tc.gregorian.Month, tc.gregorian.Day)
		if err != nil {
			t.Fatalf("GregorianToJulian(%+v): %v", tc.gregorian, err)
		}
		if gotJulian != tc.julian {
			t.Errorf("GregorianToJulian(%+v)=%+v want %+v", tc.gregorian, gotJulian, tc.julian)
		}
		gotGregorian, julianJDN, err := JulianToGregorian(tc.julian.Year, tc.julian.Month, tc.julian.Day)
		if err != nil {
			t.Fatalf("JulianToGregorian(%+v): %v", tc.julian, err)
		}
		if gotGregorian != tc.gregorian {
			t.Errorf("JulianToGregorian(%+v)=%+v want %+v", tc.julian, gotGregorian, tc.gregorian)
		}
		if gregorianJDN != julianJDN {
			t.Errorf("JDNs differ for %+v and %+v: %d != %d", tc.gregorian, tc.julian, gregorianJDN, julianJDN)
		}
	}
}

func TestCalendarSpecificLeapDayValidation(t *testing.T) {
	if _, _, err := GregorianToJulian(1900, 2, 29); err == nil {
		t.Fatal("Gregorian 1900-02-29 must be rejected")
	}
	if got, _, err := JulianToGregorian(1900, 2, 29); err != nil || got != (CivilDate{1900, 3, 13}) {
		t.Fatalf("Julian 1900-02-29 conversion = %+v, %v", got, err)
	}
	if _, _, err := JulianToGregorian(1901, 2, 29); err == nil {
		t.Fatal("Julian 1901-02-29 must be rejected")
	}
}
