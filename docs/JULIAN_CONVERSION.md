# Julian calendar conversion

“Julian date” can mean either a date in the Julian civil calendar or a Julian
Day Number used to count days. Calendar Boundary Lab now distinguishes them.

## Scope and rule

The endpoint converts between the **proleptic** Gregorian and Julian civil
calendars. Gregorian leap years use the 4/100/400 rule. Julian leap years are
every fourth year. The two calendars map to the same integer Julian Day Number
for a civil day; the astronomical Julian Date day begins at noon UT1, so
midnight is represented by a half-day fraction.

The converter does not model a country's historical adoption date. For example,
the Naval Observatory's October 1582 reform uses Julian 1582-10-04 followed by
Gregorian 1582-10-15 in Roman Catholic territories, while England and its
colonies adopted later. This API instead uses each calendar's rules
proleptically and does not skip locally omitted dates.

## Examples

```sh
curl 'http://localhost:8080/v1/convert/gregorian/2027/1/1'
curl 'http://localhost:8080/v1/convert/julian/1900/2/29'
```

The first maps Gregorian 2027-01-01 to Julian 2026-12-19. The second accepts a
Julian leap day that did not exist in the Gregorian calendar for 1900. Responses
include the source date, target date, calendar names, and the integer JDN at
noon. A civil date near year 1 may be rejected if its counterpart falls outside
the supported positive 1–9999 range.

## Tests

The conversion tests cover 2027, the proleptic 1582 alignment, distinct leap
rules in 1900 and 2000, equal day numbers for corresponding dates, and invalid
dates. The source algorithm is based on the Naval Observatory's published
calendar-to-Julian-date formulas.

See [USNO calendar conversion formulas](https://aa.usno.navy.mil/faq/JD_formula),
[USNO calendar overview](https://aa.usno.navy.mil/faq/calendars), and the
[endpoint reference](API.md).
