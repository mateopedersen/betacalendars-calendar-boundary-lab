# Calendar Invariants

- Every in-month civil date occurs exactly once.
- In-month dates increase by exactly one Gregorian day.
- The computed month row count is 4–6 and equals `ceil((leadingCells + daysInMonth) / 7)`.
- Monday-first and Sunday-first modes change only grid offsets and row count, not the civil dates.
- Leap year iff divisible by 400, or divisible by 4 and not by 100.
- ISO week-year may differ from calendar year near January 1 and December 31.
- Calendar arithmetic uses date-only values and does not depend on a machine's local timezone.
- Julian and Gregorian leap rules are checked separately; matching civil days map to the same integer Julian Day Number.
- The converter is proleptic and does not infer a country's historical Gregorian reform date.
