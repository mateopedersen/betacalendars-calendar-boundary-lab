# Architecture

The standard-library HTTP server separates civil-date arithmetic from transport. `internal/calendar` computes Gregorian month grids, row counts, year matrices, and ISO week boundaries. `internal/httpapi` exposes JSON, health probes, metrics, the fixture response, and a server-rendered engineering dashboard. It has no Kubernetes API client and needs no RBAC.

The container is built as a static Go binary in a build stage and copied into a minimal non-root runtime. Helm configures the image, service, optional ingress, HPA, PDB, network policy, and optional ServiceMonitor.
