# Calendar data model

## Civil dates are not instants

A civil date such as `2027-01-01` identifies a calendar day. It has no time of
day, UTC offset, or timezone. The service uses Go `time.Date` values at UTC
midnight as an implementation detail for Gregorian weekday and ISO-week
arithmetic; the API does not convert a user event between zones.

An instant is different: `2027-01-01T09:00:00Z` identifies a point on the UTC
timeline. A local appointment needs both a wall-clock value and a named zone,
such as `Europe/Istanbul`, if future timezone-rule changes must be applied.
An offset alone does not encode the rules for a geographic zone.

## Month response

The `/v1/month` response carries the requested Gregorian `year` and `month`, a
month name, `daysInMonth`, the first weekday, selected `weekStart`, number of
leading cells, natural row count, date rows, and three invariants. Each date row
contains seven positions. Dates inside the target month are ISO-format strings;
outside positions are empty strings.

## Year matrix and ISO fields

Each year-matrix month includes calendar geometry plus the ISO week-year and
week number of its first and last civil dates. These are separate integer
fields: January 1 can belong to the previous ISO week-year. The distinction is
covered by the 2027 January test.

## Serialization contract

JSON uses four-digit positive Gregorian years for month and date values. API
consumers should parse numeric fields as integers and should not infer an
adjacent date from an empty cell. The response contains only target-month date
strings; MonthGrid is the companion project that preserves adjacent dates for
print and browser grids.
