# Troubleshooting

## HTTP 400 from a month endpoint

Check that year is 1–9999, month is 1–12, and `weekStart` is `monday` or
`sunday`. For an unsupported value, the server returns a readable error.

## Unexpected Monday/Sunday layout

The API defaults to Monday-first. Pass `?weekStart=sunday` to override it.
`WEEK_START` is used only when the query parameter is empty.

## ISO year differs from the month label

This is expected near New Year's Day. Read the ISO week-year and ISO week
number as separate values. January 1, 2027 is in 2026-W53.

## Gregorian/Julian conversion rejected a date

Validate the date under the source calendar's own leap rule. Julian February 29
is valid every fourth Julian year, while Gregorian century years require
divisibility by 400. The converter is proleptic and does not apply a country's
historical adoption gap.

## Documentation build fails

Run `python -m pip install -r docs/requirements.txt` and
`mkdocs build --strict`. The strict build treats unresolved navigation and
internal document links as errors; external-site availability is audited
separately because it can vary.
