# Sources and audit notes

## Primary technical sources

- [U.S. Naval Observatory: introduction to calendars](https://aa.usno.navy.mil/faq/calendars) — Gregorian and Julian leap rules and calendar context.
- [U.S. Naval Observatory: Gregorian/Julian date conversion](https://aa.usno.navy.mil/faq/JD_formula) — published conversion algorithms and Julian Day Number description.
- [U.S. Naval Observatory: Julian Date converter notes](https://aa.usno.navy.mil/data/JulianDate) — noon-based Julian dates and local historical reform caveats.
- [Python `datetime` reference](https://docs.python.org/3.12/library/datetime.html) — proleptic Gregorian date model and ISO week interface.
- [Python `calendar` reference](https://docs.python.org/3.12/library/calendar.html) — proleptic Gregorian month computations.
- [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339) — Internet timestamp profile and offset syntax.
- [IANA Time Zone Database](https://www.iana.org/time-zones) and [theory and pragmatics](https://www.iana.org/time-zones/theory) — changing civil-time rules and zone data scope.
- [Read the Docs configuration v2 reference](https://docs.readthedocs.com/platform/stable/config-file/v2.html) — repository build configuration.

## BetaCalendars pages

The site was checked over HTTPS. The home page, monthly collection, blank page,
weekly page, and January–September 2027 monthly pages responded successfully
and declared indexable canonical URLs during this audit. The monthly collection
showed October–December 2026 followed by January–September 2027. No verified
2027 annual page or October–December 2027 monthly page was found, so this
handbook does not assert those resources exist. The sitemap index itself returns
`X-Robots-Tag: noindex`, as is typical for sitemap files; that header is not the
indexing policy for the HTML pages linked from it.

Canonical and robots behavior can change with the publisher's site settings.
This note records the checked state, not a promise about future crawler results.
