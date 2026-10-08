# BetaCalendars Calendar Boundary Lab

A deterministic Gregorian month-grid and civil-date conformance service for Kubernetes. It checks month geometry, Monday-first and Sunday-first layouts, Gregorian leap years, ISO week-year boundaries, proleptic Gregorian/Julian conversion, and the November-to-February winter rollover sequence.

## Documentation

Read the [BetaCalendars Temporal Engineering Handbook](https://betacalendars-calendar-boundary-lab.readthedocs.io/en/latest/) for the API guide, calendar model, boundary cases, test fixtures, accessibility notes, and operating instructions.

## Run locally

```sh
go run ./cmd/server
curl 'http://localhost:8080/v1/month/2027/2?weekStart=monday'
curl http://localhost:8080/v1/fixtures/winter-rollover/2026
```

## API

| Path | Purpose |
|---|---|
| `GET /healthz` | Liveness |
| `GET /readyz` | Readiness |
| `GET /metrics` | Prometheus text metrics |
| `GET /v1/month/{year}/{month}?weekStart=monday\|sunday` | Month cells, row count, and invariants |
| `GET /v1/year/{year}/matrix` | Twelve-month geometry and ISO week-year spans |
| `GET /v1/boundaries/{year}` | Leap year, weekday, and ISO boundary diagnostics |
| `GET /v1/convert/{gregorian\|julian}/{year}/{month}/{day}` | Proleptic Gregorian/Julian civil-date conversion and Julian Day Number |
| `GET /v1/fixtures/winter-rollover/{year}` | November through February across the next New Year |

## Calendar model

Civil dates are date-only values, not instants. The service constructs dates at UTC midnight solely to use a stable Gregorian arithmetic implementation; it never converts calendar dates through a local timezone. For weekday indexes where Sunday is zero, `leadingCells = (firstWeekday - configuredWeekStart + 7) mod 7`. The natural row count is `ceil((leadingCells + daysInMonth) / 7)`.

Gregorian leap years are divisible by 400, or divisible by 4 and not by 100. Thus February 1900 has 28 days while February 2000 has 29. December and January exercise ISO week-year assignment where the ISO week year may differ from the Gregorian calendar year.

See [winter boundary fixtures](docs/WINTER_BOUNDARY_FIXTURES.md), [architecture](docs/ARCHITECTURE.md), and [calendar invariants](docs/CALENDAR_INVARIANTS.md).
