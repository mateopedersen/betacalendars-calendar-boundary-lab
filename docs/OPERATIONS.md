# Reproducible operations

The project runs as a Go HTTP server with no Kubernetes API client and no
runtime dependency on Helm. The container uses a static build stage and a
minimal non-root runtime; the Helm chart configures deployment, service, health
probes, optional ingress, autoscaling, network policy, and Prometheus
ServiceMonitor support.

## Probes and metrics

- `/healthz` returns `ok` for liveness.
- `/readyz` returns `ready` for readiness.
- `/metrics` exposes request counters and the build version in Prometheus text
  format.

Helm defaults include one replica, a 32 MiB memory request, a 128 MiB memory
limit, a non-root UID, dropped Linux capabilities, and a read-only root
filesystem. Inspect the current chart values before deployment because chart
defaults can change over time.

## Build reproducibility

Go module dependencies are recorded in `go.mod` and `go.sum`; the service uses
the standard library for date arithmetic and HTTP handling. Documentation uses
the exact MkDocs and Material versions pinned in `docs/requirements.txt`.
Continuous integration runs Go vet/tests, Helm validation, and a strict MkDocs
build on pushes to `main` and pull requests.
