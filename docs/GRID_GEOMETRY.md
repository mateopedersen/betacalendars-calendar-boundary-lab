# Month grid geometry

For a selected first weekday, calculate the leading offset before the first of
the month. Let `weekday` use Sunday as zero and `start` be zero for Sunday-first
or one for Monday-first:

```text
leading = (weekday(YYYY-MM-01) - start + 7) mod 7
rows = ceil((leading + daysInMonth) / 7)
```

`Geometry` returns only in-month date strings, padded with empty strings before
day 1 and after the last day. This makes the API easy to consume for diagnostics
but does not supply dates from neighboring months. MonthGrid uses real adjacent
civil dates for printable grids and separately flags whether each cell belongs
to the selected month.

## Week start is a layout rule

Changing Monday-first to Sunday-first changes the leading offset and can change
the natural row count. It must not change the underlying month dates. February
2027 needs four Monday-first rows and five Sunday-first rows. May 2027 needs six
rows for both starts. Both patterns appear in the Go test vectors.

The API defaults to Monday-first; callers can select `?weekStart=sunday`. A
configured `WEEK_START` environment value is used only when the query parameter
is empty.
