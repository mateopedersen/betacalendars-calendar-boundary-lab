# Accessibility notes

The dashboard uses a document title, a main heading, section headings, native
links, a native `select`, and responsive CSS that changes its card layout on
narrow screens. JSON endpoints can be consumed without relying on dashboard
colors or visual status pills.

The current project does not claim a complete WCAG audit. Its dashboard has no
automated screen-reader, keyboard, zoom, or contrast regression suite. Before
using it as a public-facing calendar UI, add those checks and ensure status is
not communicated by color alone. The underlying date API is deliberately
presentation-neutral.

MonthGrid's separate HTML renderer uses a captioned semantic table, column
headers, `<time datetime>` values, and an accessible label for adjacent dates.
Its SVG includes a title and description. See the `docs/accessibility.md` page
in the [MonthGrid repository](https://git.sr.ht/~mateopedersen/monthgrid).
