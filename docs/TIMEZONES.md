# Time zones and daylight-saving transitions

## Keep a month grid date-only

Month and weekday arithmetic should operate on a civil date, not by repeatedly
adding 24 hours to a timestamp. Local days around offset transitions can be
shorter or longer than 24 elapsed hours. Calendar Boundary Lab constructs UTC
midnight values solely to use Go's date and ISO-week arithmetic; it does not
store appointment times or derive local event dates.

For timed events, persist an unambiguous timestamp with an offset and the
geographic timezone identifier when future local rules matter. RFC 3339 defines
an Internet timestamp profile with a numeric offset or `Z`; an offset does not
carry a location's future rule set. IANA's timezone database records political
changes to offsets and daylight-saving rules and is updated over time.

## Recommended separation

1. Store a calendar date as year, month, and day.
2. Store an event instant as a UTC timestamp.
3. Store the intended local zone as an IANA name such as `Europe/Istanbul`.
4. Convert between them only at the application boundary, using an updated
   timezone database and explicit ambiguity policy.

This service does not include a scheduling/timezone API. See the primary
references: [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339),
[IANA Time Zone Database](https://www.iana.org/time-zones),
[IANA theory](https://www.iana.org/time-zones/theory), and
[Python zoneinfo](https://docs.python.org/3.12/library/zoneinfo.html).
