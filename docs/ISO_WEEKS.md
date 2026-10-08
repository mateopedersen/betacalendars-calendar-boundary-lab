# ISO week years

ISO week dates are not simply Gregorian dates with a week number attached.
Weeks begin on Monday. Week 1 is the week containing the first Thursday of the
Gregorian year (equivalently, January 4). As a result, the ISO week-year can be
the previous or next Gregorian year near New Year's Day.

For 2027, January 1 is Friday and belongs to ISO week 53 of 2026; Monday,
January 4 begins ISO week 1 of 2027. December 31, 2027 belongs to ISO week 52
of 2027. These facts are computed with Go's `time.Time.ISOWeek()` and exposed
separately as the first and last ISO week-year/week fields in
`/v1/year/2027/matrix`.

```sh
curl 'http://localhost:8080/v1/boundaries/2027'
```

Use ISO week-year plus week number as a pair. A week label such as `2026-W53`
must not be rebuilt from a date's Gregorian year. MonthGrid preserves ISO
metadata per cell in its JSON fixtures and includes the corresponding January
2027 case study in the [MonthGrid repository](https://git.sr.ht/~mateopedersen/monthgrid).

The formal ISO 8601 text is paywalled; the implementation uses the standard
library's ISO calendar function. See [Python's date ISO calendar reference](https://docs.python.org/3.12/library/datetime.html#datetime.date.isocalendar)
for an accessible description of the same fields and week-year boundary.
