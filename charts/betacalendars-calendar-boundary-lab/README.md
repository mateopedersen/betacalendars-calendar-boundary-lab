# BetaCalendars Calendar Boundary Lab

## What It Does

Calendar Boundary Lab is a deterministic Gregorian calendar boundary, month-grid, and civil-date conformance service for Kubernetes. It provides JSON APIs and an engineering diagnostics dashboard for validating weekdays, month rows, leap years, year rollover, and ISO week-year edges. It does not scrape BetaCalendars or use it as an API.

## Why Calendar Boundaries Are Hard

Calendar dates are not timestamps. A date such as `2027-02-01` has no time or timezone; converting it through local time can shift the displayed day. The service uses UTC only as an arithmetic convention and returns date-only values. Gregorian leap years are divisible by 400, or divisible by 4 but not 100. February 1900 is therefore 28 days; February 2000 is 29.

For weekday values Sunday=0 through Saturday=6, `leadingCells = (firstWeekday - configuredWeekStart + 7) mod 7`. The number of natural rows is `ceil((leadingCells + daysInMonth) / 7)`. December and January are especially useful because ISO week numbering can assign early January to the previous ISO week-year or late December to the next one.

## Architecture

```text
HTTP handlers ──> civil-date arithmetic ──> month geometry
      │                     │                    ├── Monday/Sunday policies
      │                     ├── leap-year rules  └── row/date invariants
      ├── winter rollover fixture suite
      ├── Prometheus metrics
      └── server-rendered diagnostics UI
```

## API

| Endpoint | Response |
|---|---|
| `GET /healthz` | Liveness |
| `GET /readyz` | Readiness |
| `GET /metrics` | Prometheus metrics |
| `GET /v1/month/{year}/{month}?weekStart=monday` | Month geometry; `weekStart` also accepts `sunday` |
| `GET /v1/year/{year}/matrix` | Twelve month geometries and ISO week-year start/end details |
| `GET /v1/boundaries/{year}` | Gregorian and ISO boundary checks |
| `GET /v1/fixtures/winter-rollover/{year}` | November and December of the requested year, then January and February of the next year |

## Winter Rollover Conformance Suite

The November → December → January → February sequence covers approach to year-end, Gregorian rollover, ISO week-year divergence, new-year alignment, and February's 28/29-day behavior. Four months also expose row-count changes under both Monday-first and Sunday-first layouts. Algorithmic tests assert date arithmetic and invariants; visual fixtures help a person compare rendered grids. These methods catch different classes of errors.

Manual visual fixtures: [November calendar visual fixture](https://www.betacalendars.com/november-calendar.html), [December calendar visual fixture](https://www.betacalendars.com/december-calendar.html), [January calendar visual fixture](https://www.betacalendars.com/january-calendar.html), and [February calendar visual fixture](https://www.betacalendars.com/february-calendar.html). These are optional human-readable fixtures for manual comparison. They are not application dependencies. [Beta Calendars](https://www.betacalendars.com/)

## Installation

```sh
helm repo add betacalendars https://mateopedersen.github.io/betacalendars-calendar-boundary-lab/
helm repo update
helm install boundary-lab betacalendars/betacalendars-calendar-boundary-lab
```

## Configuration

| Value | Default | Description |
|---|---:|---|
| `replicaCount` | `1` | Fixed replicas when autoscaling is disabled |
| `image.repository` | `ghcr.io/mateopedersen/betacalendars-calendar-boundary-lab` | Public GHCR image repository |
| `image.tag` | `1.0.0` | Application image tag |
| `service.port` | `8080` | HTTP service port |
| `calendar.defaultYear` | `2027` | Dashboard year |
| `calendar.weekStart` | `monday` | Default layout (`monday` or `sunday`) |
| `ingress.enabled` | `false` | Create a networking.k8s.io/v1 Ingress |
| `autoscaling.enabled` | `false` | Create an autoscaling/v2 HPA |
| `networkPolicy.enabled` | `false` | Create an ingress-only NetworkPolicy |
| `serviceMonitor.enabled` | `false` | Create a Prometheus Operator ServiceMonitor |

## Hardened Installation Example

```sh
helm install boundary-lab betacalendars/betacalendars-calendar-boundary-lab \\
  --set podSecurityContext.runAsNonRoot=true \\
  --set securityContext.readOnlyRootFilesystem=true \\
  --set securityContext.allowPrivilegeEscalation=false
```

The chart runs as UID/GID 10001, disables privilege escalation, uses a read-only root filesystem, drops all Linux capabilities, sets RuntimeDefault seccomp, and does not mount a Kubernetes API token. The service needs no Kubernetes API permissions.

## Prometheus Metrics

`calendar_requests_total` and `calendar_fixture_runs_total` expose request and winter-fixture counters. ServiceMonitor creation is optional and disabled by default.

## Network Policy

NetworkPolicy is off by default. Enable it only with a cluster network plugin that enforces policies, and set permitted ingress sources for the service.

## Ingress

Ingress is disabled by default. When enabled, the chart emits `networking.k8s.io/v1` rules and accepts host, path, annotation, and TLS values.

## Horizontal Pod Autoscaling

Autoscaling is disabled by default. Enabling it creates an `autoscaling/v2` CPU HPA; the cluster must provide metrics-server and the chart omits fixed Deployment replicas.

## Testing

Run `go test ./...`, `go vet ./...`, `helm lint charts/betacalendars-calendar-boundary-lab`, and render default and hardened values with `helm template`.

## Security

See the repository [SECURITY.md](../../SECURITY.md). No credentials or application data are required. The optional visual references are never fetched by the service.

## Calendar Model

Weekday and month-grid calculations use the proleptic Gregorian calendar provided by Go's standard library. Dates are date-only and never timezone-converted. Supported year inputs are 1–9998 for year-boundary operations.

## Limitations

This service validates Gregorian civil-date geometry only. It does not model locale-specific holidays, historical calendar reforms, event recurrence, timezone rules, or time-of-day behavior.

## License

Apache-2.0.
