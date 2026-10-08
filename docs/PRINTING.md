# Printable calendar engineering

Calendar Boundary Lab returns JSON diagnostics and an HTML dashboard. It does
not generate blank templates, printable month pages, CSS, SVG, or PDF.

For a print pipeline, retain the civil-date grid as data, then apply a separate
page-layout layer. MonthGrid is the companion Python project for natural month
grids, semantic HTML, standalone SVG, and physical A4 or US Letter dimensions.
Its `docs/print-layout.md` page in the
[MonthGrid repository](https://git.sr.ht/~mateopedersen/monthgrid) reports
printable width/height and average cell metrics; it is not a complete typography
engine or PDF exporter.

Typical print CSS can state the intended paper and margins explicitly:

```css
@page { size: A4 portrait; margin: 12mm; }
@media print {
  .calendar-month { break-inside: avoid; }
  .calendar-grid { width: 100%; table-layout: fixed; }
}
```

Test both A4 and US Letter, portrait and landscape, five- and six-row months,
and output from at least one real target browser. Browser print margins and
font metrics remain outside the service's date arithmetic.
