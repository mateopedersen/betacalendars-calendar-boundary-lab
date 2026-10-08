package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	cal "github.com/betacalendars/betacalendars-calendar-boundary-lab/internal/calendar"
)

func TestMonthEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/v1/month/2027/2?weekStart=monday", nil)
	w := httptest.NewRecorder()
	New().ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}
	var got map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["daysInMonth"] != float64(28) || got["naturalRows"] != float64(4) {
		t.Fatalf("unexpected response: %#v", got)
	}
}

func TestWinterRolloverEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/v1/fixtures/winter-rollover/2026", nil)
	w := httptest.NewRecorder()
	New().ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}
	var got struct {
		Months []struct {
			Year  int `json:"year"`
			Month int `json:"month"`
		} `json:"months"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	want := [][2]int{{2026, 11}, {2026, 12}, {2027, 1}, {2027, 2}}
	if len(got.Months) != len(want) {
		t.Fatalf("got %d fixture months", len(got.Months))
	}
	for i, m := range got.Months {
		if m.Year != want[i][0] || m.Month != want[i][1] {
			t.Errorf("fixture %d=%d/%d want %d/%d", i, m.Year, m.Month, want[i][0], want[i][1])
		}
	}
}

func TestProlepticGregorianJulianConversionEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/v1/convert/gregorian/2027/1/1", nil)
	w := httptest.NewRecorder()
	New().ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}
	var got struct {
		SourceCalendar        string `json:"sourceCalendar"`
		TargetCalendar        string `json:"targetCalendar"`
		Source                struct{ Year, Month, Day int }
		Target                struct{ Year, Month, Day int }
		JulianDayNumberAtNoon int `json:"julianDayNumberAtNoon"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.SourceCalendar != "gregorian" || got.TargetCalendar != "julian" {
		t.Fatalf("unexpected calendars: %#v", got)
	}
	if got.Target.Year != 2026 || got.Target.Month != 12 || got.Target.Day != 19 {
		t.Fatalf("unexpected converted date: %#v", got.Target)
	}
	if got.JulianDayNumberAtNoon != 2461407 {
		t.Fatalf("unexpected JDN: %d", got.JulianDayNumberAtNoon)
	}
}

func TestYearMatrixReportsISOWeekYearAndWeekSeparately(t *testing.T) {
	months, err := cal.YearMatrix(2027)
	if err != nil {
		t.Fatal(err)
	}
	if got := months[0]; got.ISOStartYear != 2026 || got.ISOStartWeek != 53 || got.ISOEndYear != 2027 || got.ISOEndWeek != 4 {
		t.Fatalf("January ISO span=%#v", got)
	}
}
