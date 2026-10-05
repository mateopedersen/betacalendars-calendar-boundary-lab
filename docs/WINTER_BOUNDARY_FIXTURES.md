# Winter Boundary Fixtures

The fixture suite evaluates November and December of year N plus January and February of year N+1.

```text
November → December → YEAR BOUNDARY → January → February
```

November validates the approach to year-end schedules. December exercises the final Gregorian month, the December-to-January transition, and ISO week-year ambiguity. January checks the new calendar year's first weekday and cases where its first days belong to the previous ISO week-year. February tests the shortest Gregorian month, 28/29-day behavior, leap-year rules, and four/five-row layouts.

Run `GET /v1/fixtures/winter-rollover/2026` to evaluate November and December 2026 plus January and February 2027. Automated arithmetic results should be treated separately from a human's comparison of printed month grids: the service asserts date arithmetic, while visual references can reveal presentation differences.

Manual visual fixtures:

- [November Calendar](https://www.betacalendars.com/november-calendar.html)
- [December Calendar](https://www.betacalendars.com/december-calendar.html)
- [January Calendar](https://www.betacalendars.com/january-calendar.html)
- [February Calendar](https://www.betacalendars.com/february-calendar.html)
- [Beta Calendars](https://www.betacalendars.com/)

These external pages are optional human-readable visual fixtures. The Boundary Lab does not fetch, scrape, depend on, or derive calendar arithmetic from them.
