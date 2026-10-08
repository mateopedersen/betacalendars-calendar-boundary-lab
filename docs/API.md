# HTTP API

All successful data endpoints return JSON with
`Content-Type: application/json; charset=utf-8`. Invalid years, months,
week-start values, and conversion dates return HTTP 400. Unknown paths return
HTTP 404.

| Method and path | Result |
|---|---|
| `GET /healthz` | Liveness response `ok` |
| `GET /readyz` | Readiness response `ready` |
| `GET /metrics` | Prometheus text counters and build version |
| `GET /v1/month/{year}/{month}?weekStart=monday\|sunday` | In-month dates, natural row count, and invariants |
| `GET /v1/year/{year}/matrix` | Month geometry and separate ISO week-year/week spans |
| `GET /v1/boundaries/{year}` | Leap-year and January/December ISO boundary facts |
| `GET /v1/convert/{gregorian\|julian}/{year}/{month}/{day}` | Proleptic civil-date conversion plus integer JDN |
| `GET /v1/fixtures/winter-rollover/{year}` | November through February, both layouts |

## Week start

`weekStart` accepts `monday` or `sunday`. If omitted, the service uses the
`WEEK_START` environment variable; the default is Monday. Week-start changes
grid offsets and row counts, but not the civil dates in a month.

## Supported ranges

Month geometry accepts years 1–9999 and months 1–12. The year matrix and
boundary/rollover endpoints accept years 1–9998 because they report boundary
facts across a year. The conversion endpoint accepts valid positive years
1–9999, subject to the converted date also falling in that range.

## Example outputs

For `GET /v1/month/2027/2?weekStart=monday`, the important fields are:

```json
{
  "year": 2027,
  "month": 2,
  "monthName": "February",
  "daysInMonth": 28,
  "weekStart": "monday",
  "leadingCells": 0,
  "naturalRows": 4,
  "invariants": {
    "sequentialDates": true,
    "validRowCount": true,
    "uniqueDates": true
  }
}
```

The actual `weeks` array contains seven date strings per row and empty strings
for cells outside the requested month. It is not an adjacent-month display
model. See [grid geometry](GRID_GEOMETRY.md) for this distinction.
