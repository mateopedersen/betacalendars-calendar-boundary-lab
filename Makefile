.PHONY: test vet build helm-lint helm-template

test:
	go test ./...

vet:
	go vet ./...

build:
	CGO_ENABLED=0 go build -trimpath -o dist/boundary-lab ./cmd/server

helm-lint:
	helm lint charts/betacalendars-calendar-boundary-lab

helm-template:
	helm template boundary-lab charts/betacalendars-calendar-boundary-lab
