# Fixtures and automated tests

Run the same checks used by CI:

```sh
go vet ./...
go test ./...
mkdocs build --strict
python scripts/check_external_docs_links.py
```

The Go tests cover Gregorian century leap rules; month lengths, weekdays, and
natural row counts for all 2027 months; date uniqueness and sequence under both
week starts; January's ISO year boundary; HTTP response behavior; winter
rollover payloads; proleptic Gregorian/Julian conversion; and invalid leap-day
inputs.

The winter fixture endpoint takes a base year `N` and returns November and
December of `N`, then January and February of `N+1`. It validates civil date
geometry; visual comparison is a separate manual step. Beta Calendars links in
the fixture document are references for a human checking presentation, not an
automated source of expected dates.

The tests assert the algorithm's date values and row counts, not pixel-perfect
dashboard rendering, timezone database behavior, browser print output, or
historical local calendar adoptions. The companion
[MonthGrid repository](https://git.sr.ht/~mateopedersen/monthgrid) has its own
month-by-month JSON/HTML/SVG fixtures and A4/Letter geometry tests.
