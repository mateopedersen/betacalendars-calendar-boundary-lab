# Architecture

The standard-library HTTP server separates civil-date arithmetic from transport. `internal/calendar` computes proleptic Gregorian and Julian conversions, Gregorian month grids, row counts, year matrices, and ISO week boundaries. `internal/httpapi` exposes JSON, health probes, metrics, month and winter fixtures, and a server-rendered engineering dashboard. It has no Kubernetes API client and needs no RBAC.

The container is built as a static Go binary in a build stage and copied into a minimal non-root runtime. Helm configures the image, service, optional ingress, HPA, PDB, network policy, and optional ServiceMonitor.

The civil-date conversion routines are kept in `internal/calendar/conversion.go`
and operate proleptically. They do not apply local historical reform dates.
Read the [API reference](API.md) and [Julian conversion notes](JULIAN_CONVERSION.md)
for the interface and limitations.
