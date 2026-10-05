package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMonthEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/v1/month/2027/2?weekStart=monday", nil)
	w := httptest.NewRecorder()
	New().ServeHTTP(w, r)
	if w.Code != http.StatusOK { t.Fatalf("status=%d body=%s", w.Code, w.Body.String()) }
	var got map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil { t.Fatal(err) }
	if got["daysInMonth"] != float64(28) || got["naturalRows"] != float64(4) { t.Fatalf("unexpected response: %#v", got) }
}

func TestWinterRolloverEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/v1/fixtures/winter-rollover/2026", nil)
	w := httptest.NewRecorder()
	New().ServeHTTP(w, r)
	if w.Code != http.StatusOK { t.Fatalf("status=%d body=%s", w.Code, w.Body.String()) }
	var got struct { Months []struct { Year int `json:"year"`; Month int `json:"month"` } `json:"months"` }
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil { t.Fatal(err) }
	want := [][2]int{{2026,11},{2026,12},{2027,1},{2027,2}}
	if len(got.Months) != len(want) { t.Fatalf("got %d fixture months",len(got.Months)) }
	for i,m := range got.Months { if m.Year != want[i][0] || m.Month != want[i][1] { t.Errorf("fixture %d=%d/%d want %d/%d",i,m.Year,m.Month,want[i][0],want[i][1]) } }
}
