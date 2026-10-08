# BetaCalendars Temporal Engineering Handbook

This handbook explains the date arithmetic implemented by Calendar Boundary
Lab, the assumptions around it, and the interfaces that expose its results.
Examples are tied to the Go source and test suite in the public repository.

## What the project computes

Calendar Boundary Lab produces proleptic Gregorian month geometry, Monday- or
Sunday-first row layouts, year matrices, leap-year and ISO-week diagnostics,
November-to-February rollover fixtures, and proleptic Gregorian/Julian civil
date conversions. It is a deterministic Go HTTP service with a server-rendered
diagnostic dashboard, a container image, and a Helm chart.

It does not provide holidays, locale-specific translations, timezone-aware
events, a PDF generator, a blank-calendar endpoint, or a general calendar
publishing API. Those capabilities belong in separate layers or tools.

## Read in order

1. [Quickstart](QUICKSTART.md) — run the tests and query the local service.
2. [Data model](DATA_MODEL.md) — distinguish civil dates, timestamps, and ISO week years.
3. [Gregorian arithmetic](GREGORIAN.md) and [Julian conversion](JULIAN_CONVERSION.md).
4. [ISO weeks](ISO_WEEKS.md), [grid geometry](GRID_GEOMETRY.md), and
   [time-zone boundaries](TIMEZONES.md).
5. [Fixtures and testing](TESTING.md) — reproduce the edge cases in CI.

## Source and references

- [Calendar Boundary Lab source](https://github.com/mateopedersen/betacalendars-calendar-boundary-lab)
- [MonthGrid Python renderer and print geometry](https://git.sr.ht/~mateopedersen/monthgrid)
- [Beta Calendars home](https://www.betacalendars.com/)
- [Beta Calendars monthly collection](https://www.betacalendars.com/monthly-calendar)

The Beta Calendars links are visual references, not API dependencies. The lab
does not fetch their pages or use them to calculate dates.
