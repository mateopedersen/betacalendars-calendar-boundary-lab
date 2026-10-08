# Quickstart

## Requirements and tests

The service uses the Go standard library. With Go 1.25 or newer installed:

```sh
go test ./...
go vet ./...
go run ./cmd/server
```

The server listens on port 8080 by default. Set `PORT` to choose another valid
port. Open `http://localhost:8080/` for the diagnostics dashboard.

## Query an API example

In another terminal, request February 2027 with Monday-first columns:

```sh
curl 'http://localhost:8080/v1/month/2027/2?weekStart=monday'
```

The JSON response reports 28 dates and four natural rows. Query the full year's
geometry and ISO spans with:

```sh
curl 'http://localhost:8080/v1/year/2027/matrix'
```

The year matrix is diagnostic data, not a rendered or printable twelve-month
calendar. For a winter transition, request:

```sh
curl 'http://localhost:8080/v1/fixtures/winter-rollover/2026'
```

That response contains November and December 2026 and January and February
2027, each in both week-start layouts.

## Run without a container

`go run ./cmd/server` compiles and starts the same standard-library server used
in the container. It does not require Kubernetes or credentials. See the
[API reference](API.md) before relying on response fields in another service.
