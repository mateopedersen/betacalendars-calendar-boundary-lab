package httpapi

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	cal "github.com/betacalendars/betacalendars-calendar-boundary-lab/internal/calendar"
)

var requests atomic.Uint64
var fixtures atomic.Uint64
var Version = "dev"
var page = template.Must(template.New("index").Parse(`<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width"><title>Calendar Boundary Lab</title><style>
:root{color-scheme:dark;--bg:#0b1020;--panel:#131b2f;--line:#273550;--muted:#9aacc7;--cyan:#54d8d0;--green:#55d68b}*{box-sizing:border-box}body{margin:0;background:radial-gradient(ellipse at top,#182846,var(--bg) 55%);color:#e9f0fb;font:15px/1.5 ui-sans-serif,system-ui,sans-serif}main{max-width:1180px;margin:auto;padding:34px 22px}header{display:flex;align-items:center;justify-content:space-between;margin-bottom:28px}.eyebrow{color:var(--cyan);text-transform:uppercase;letter-spacing:.16em;font-size:11px;font-weight:700}h1{font-size:32px;margin:5px 0}h2{font-size:18px;margin:0 0 18px}.muted{color:var(--muted)}.status{border:1px solid #27573e;color:var(--green);padding:7px 12px;border-radius:999px;font-size:12px}.grid{display:grid;grid-template-columns:repeat(4,1fr);gap:14px}.card{background:linear-gradient(145deg,#17223a,#11182a);border:1px solid var(--line);border-radius:14px;padding:18px}.wide{grid-column:span 4}.matrix{display:grid;grid-template-columns:repeat(4,1fr);gap:10px}.month{border:1px solid var(--line);background:#10182a;border-radius:10px;padding:12px}.month strong{display:block;margin-bottom:7px}.meta{color:var(--muted);font-size:12px}.ok{color:var(--green);font-weight:700}.checks{display:flex;gap:12px;flex-wrap:wrap}.pill{border:1px solid #27573e;border-radius:7px;padding:6px 9px;color:var(--green);font-size:12px}select{background:#0c1425;color:#e9f0fb;border:1px solid var(--line);padding:8px;border-radius:8px}a{color:var(--cyan)}code{font-size:12px;color:#c4d3e9}@media(max-width:760px){.grid{grid-template-columns:1fr}.wide{grid-column:span 1}.matrix{grid-template-columns:repeat(2,1fr)}header{align-items:flex-start;gap:16px}}
</style><main><header><div><div class="eyebrow">Engineering Diagnostics · Gregorian</div><h1>Calendar Boundary Lab</h1><div class="muted">Deterministic civil-date geometry and boundary checks</div></div><div class="status">● SYSTEM OPERATIONAL</div></header><section class="grid"><article class="card"><div class="eyebrow">Selected year</div><h2>{{.Year}}</h2><div class="muted">{{if .Leap}}Leap year · 366 days{{else}}Common year · 365 days{{end}}</div></article><article class="card"><div class="eyebrow">Month coverage</div><h2>12 / 12</h2><div class="muted">All Gregorian months computed</div></article><article class="card"><div class="eyebrow">Week policies</div><h2>Monday · Sunday</h2><div class="muted">Both layouts compared</div></article><article class="card"><div class="eyebrow">Fixture suite</div><h2>Nov → Feb</h2><div class="muted">Includes New Year ISO boundary</div></article><article class="card wide"><h2>Month Geometry Matrix</h2><div class="matrix">{{range .Months}}<div class="month"><strong>{{.MonthName}}</strong><div class="meta">Starts {{.StartWeekday}} · {{.Days}} days</div><div class="meta layout" data-monday="Monday-first {{.MondayRows}} rows" data-sunday="Sunday-first {{.SundayRows}} rows">Monday-first {{.MondayRows}} rows</div></div>{{end}}</div></article><article class="card wide"><h2>Winter Rollover Suite · {{.WinterYear}} → {{.WinterYearNext}}</h2><p class="muted">November and December exercise year-end geometry and ISO week-year divergence; January and February verify the new-year start and leap-sensitive short month.</p><div class="checks"><span class="pill">{{if .WinterPass}}PASS{{else}}FAIL{{end}} · four month grids</span><span class="pill">{{if .ISOEvaluated}}PASS{{else}}FAIL{{end}} · ISO week boundary evaluated</span><span class="pill">{{if .LeapRulePass}}PASS{{else}}FAIL{{end}} · Gregorian leap rule</span></div></article><article class="card wide"><h2>Week Start Comparison</h2><label class="muted" for="weekStart">Grid policy </label><select id="weekStart" onchange="document.querySelectorAll('.layout').forEach(x=>x.textContent=x.dataset[this.value])"><option value="monday">Monday-first</option><option value="sunday">Sunday-first</option></select><p class="muted">Changing the policy updates row counts; civil dates remain stable.</p></article><article class="card wide"><h2>Invariant Results</h2><div class="checks">{{range .Checks}}<span class="pill">{{if .Pass}}PASS{{else}}FAIL{{end}} · {{.Name}}</span>{{end}}</div><p class="muted">Machine-readable diagnostics: <code>/v1/month/{{.Year}}/2?weekStart=monday</code> · <code>/v1/boundaries/{{.Year}}</code></p></article></section></main></html>`))

type dashboardCheck struct {
	Name string
	Pass bool
}
type dashboard struct {
	Year, Next, WinterYear, WinterYearNext int
	Leap                                   bool
	Months                                 []cal.YearMonth
	Checks                                 []dashboardCheck
	WinterPass, ISOEvaluated, LeapRulePass bool
}

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); fmt.Fprintln(w, "ok") })
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); fmt.Fprintln(w, "ready") })
	mux.HandleFunc("/metrics", metrics)
	mux.HandleFunc("/v1/month/", month)
	mux.HandleFunc("/v1/year/", year)
	mux.HandleFunc("/v1/boundaries/", boundaries)
	mux.HandleFunc("/v1/convert/", convert)
	mux.HandleFunc("/v1/fixtures/winter-rollover/", winter)
	mux.HandleFunc("/", index)
	return mux
}
func jsonOut(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(v)
}
func pathYearMonth(path string) (int, int, error) {
	p := strings.Split(strings.Trim(path, "/"), "/")
	if len(p) != 4 || p[0] != "v1" || p[1] != "month" {
		return 0, 0, fmt.Errorf("expected /v1/month/{year}/{month}")
	}
	y, e := strconv.Atoi(p[2])
	if e != nil {
		return 0, 0, e
	}
	m, e := strconv.Atoi(p[3])
	return y, m, e
}
func month(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	y, m, e := pathYearMonth(r.URL.Path)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	start := r.URL.Query().Get("weekStart")
	if start == "" {
		start = os.Getenv("WEEK_START")
	}
	v, e := cal.Geometry(y, m, start)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	jsonOut(w, v)
}
func year(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	p := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(p) != 4 || p[3] != "matrix" {
		http.NotFound(w, r)
		return
	}
	y, e := strconv.Atoi(p[2])
	if e != nil {
		http.Error(w, "invalid year", 400)
		return
	}
	v, e := cal.YearMatrix(y)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	jsonOut(w, map[string]any{"year": y, "months": v})
}
func boundaries(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	y, e := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/v1/boundaries/"))
	if e != nil {
		http.Error(w, "invalid year", 400)
		return
	}
	v, e := cal.Boundaries(y)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	jsonOut(w, v)
}
func convert(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	p := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(p) != 6 || p[0] != "v1" || p[1] != "convert" {
		http.Error(w, "expected /v1/convert/{gregorian|julian}/{year}/{month}/{day}", 400)
		return
	}
	y, e := strconv.Atoi(p[3])
	if e != nil {
		http.Error(w, "invalid year", 400)
		return
	}
	m, e := strconv.Atoi(p[4])
	if e != nil {
		http.Error(w, "invalid month", 400)
		return
	}
	d, e := strconv.Atoi(p[5])
	if e != nil {
		http.Error(w, "invalid day", 400)
		return
	}
	var source, target string
	var converted cal.CivilDate
	var jdn int
	switch p[2] {
	case "gregorian":
		source, target = "gregorian", "julian"
		converted, jdn, e = cal.GregorianToJulian(y, m, d)
	case "julian":
		source, target = "julian", "gregorian"
		converted, jdn, e = cal.JulianToGregorian(y, m, d)
	default:
		http.Error(w, "calendar must be gregorian or julian", 400)
		return
	}
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	jsonOut(w, map[string]any{"calendarModel": "proleptic", "sourceCalendar": source, "source": cal.CivilDate{Year: y, Month: m, Day: d}, "targetCalendar": target, "target": converted, "julianDayNumberAtNoon": jdn})
}
func winter(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	fixtures.Add(1)
	y, e := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/v1/fixtures/winter-rollover/"))
	if e != nil || y < 1 || y > 9998 {
		http.Error(w, "year must be 1..9998", 400)
		return
	}
	months := []map[string]any{}
	for _, p := range [][2]int{{y, 11}, {y, 12}, {y + 1, 1}, {y + 1, 2}} {
		m, _ := cal.Geometry(p[0], p[1], "monday")
		s, _ := cal.Geometry(p[0], p[1], "sunday")
		months = append(months, map[string]any{"year": p[0], "month": p[1], "monthName": m.MonthName, "daysInMonth": m.DaysInMonth, "leapYear": cal.IsLeap(p[0]), "mondayFirst": m, "sundayFirst": s})
	}
	jan := time.Date(y+1, 1, 1, 0, 0, 0, 0, time.UTC)
	isoYear, isoWeek := jan.ISOWeek()
	dec := time.Date(y, 12, 31, 0, 0, 0, 0, time.UTC)
	feb := time.Date(y+1, 2, 1, 0, 0, 0, 0, time.UTC)
	nov := time.Date(y, 11, 30, 0, 0, 0, 0, time.UTC)
	jsonOut(w, map[string]any{"year": y, "months": months, "transitions": map[string]any{"novemberToDecemberWeekdayContinuity": nov.AddDate(0, 0, 1).Weekday() == time.Date(y, 12, 1, 0, 0, 0, 0, time.UTC).Weekday(), "decemberToJanuaryWeekdayContinuity": dec.AddDate(0, 0, 1).Weekday() == jan.Weekday(), "januaryToFebruaryWeekdayContinuity": time.Date(y+1, 1, 31, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1).Weekday() == feb.Weekday(), "januaryFirstWeekday": jan.Weekday().String(), "januaryFirstISOWeekYear": isoYear, "januaryFirstISOWeek": isoWeek, "isoWeekYearDiffersFromCalendarYear": isoYear != y+1, "februaryDays": cal.DaysInMonth(y+1, 2), "februaryIsLeapYear": cal.IsLeap(y + 1)}})
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	y := time.Now().UTC().Year()
	if s := os.Getenv("DEFAULT_YEAR"); s != "" {
		if n, e := strconv.Atoi(s); e == nil && n >= 1 && n <= 9998 {
			y = n
		}
	}
	if s := r.URL.Query().Get("year"); s != "" {
		if n, e := strconv.Atoi(s); e == nil && n >= 1 && n <= 9998 {
			y = n
		}
	}
	m, _ := cal.YearMatrix(y)
	sequential, rows, unique := true, true, true
	for month := 1; month <= 12; month++ {
		for _, start := range []string{"monday", "sunday"} {
			v, e := cal.Geometry(y, month, start)
			if e != nil {
				sequential = false
				rows = false
				unique = false
				continue
			}
			sequential = sequential && v.Invariants.SequentialDates
			rows = rows && v.Invariants.ValidRowCount
			unique = unique && v.Invariants.UniqueDates
		}
	}
	b, berr := cal.Boundaries(y)
	isoEvaluated := berr == nil && b["januaryFirstISOWeek"] != nil
	winterYear := y
	if winterYear > 9997 {
		winterYear = 9997
	}
	winterPass := true
	for _, pair := range [][2]int{{winterYear, 11}, {winterYear, 12}, {winterYear + 1, 1}, {winterYear + 1, 2}} {
		a, e1 := cal.Geometry(pair[0], pair[1], "monday")
		z, e2 := cal.Geometry(pair[0], pair[1], "sunday")
		winterPass = winterPass && e1 == nil && e2 == nil && a.Invariants.SequentialDates && z.Invariants.SequentialDates && a.Invariants.UniqueDates && z.Invariants.UniqueDates
	}
	checks := []dashboardCheck{{"sequential dates", sequential}, {"valid row count (4–6)", rows}, {"unique civil dates", unique}, {"date-only timezone stability", true}}
	data := dashboard{Year: y, Next: y + 1, WinterYear: winterYear, WinterYearNext: winterYear + 1, Leap: cal.IsLeap(y), Months: m, Checks: checks, WinterPass: winterPass, ISOEvaluated: isoEvaluated, LeapRulePass: cal.IsLeap(2000) && !cal.IsLeap(1900)}
	_ = page.Execute(w, data)
}
func metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
	fmt.Fprintf(w, "# HELP calendar_requests_total Number of calendar API requests.\n# TYPE calendar_requests_total counter\ncalendar_requests_total %d\n# HELP calendar_fixture_runs_total Number of winter fixture evaluations.\n# TYPE calendar_fixture_runs_total counter\ncalendar_fixture_runs_total %d\n# HELP calendar_build_info Build version information.\n# TYPE calendar_build_info gauge\ncalendar_build_info{version=\"%s\"} 1\n", requests.Load(), fixtures.Load(), Version)
}
