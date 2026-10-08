# Gregorian arithmetic

## Leap-year rule

A year is a Gregorian leap year if it is divisible by 400, or if it is divisible
by 4 but not by 100. This is implemented by `calendar.IsLeap`. The tests include
1900 and 2100 as common centuries, and 2000 and 2400 as leap centuries.

```text
leap(y) = (y mod 400 = 0) OR ((y mod 4 = 0) AND (y mod 100 ≠ 0))
```

For month geometry, the first day of the month supplies the weekday offset and
the month length supplies the number of in-month cells. Given `leading` cells
before day 1 and `days` in the month:

```text
rows = ceil((leading + days) / 7)
     = (leading + days + 6) integer-divided-by 7
```

The service uses the proleptic Gregorian rules for its supported year range.
That is a computational model, not a claim that every historical jurisdiction
used Gregorian dates before its local adoption.

## Verified boundaries

- February 1900 has 28 days; February 2000 has 29.
- February 2100 has 28 days; February 2400 has 29.
- Every 2027 month is tested in both Sunday-first and Monday-first layouts.
- Each in-month date appears once, in order, and row counts remain between four
  and six.

The implementation and tests are in the repository's `internal/calendar`
package. See the [source](https://github.com/mateopedersen/betacalendars-calendar-boundary-lab/tree/main/internal/calendar).
