.PHONY: generate
generate:
generate: generate/sqlc

.PHONY: generate/sqlc
generate/sqlc:
	cd riverdriver/riverdatabasesql/internal/dbsqlc && sqlc generate
	cd riverdriver/riverpgxv5/internal/dbsqlc && sqlc generate

.PHONY: lint
lint:
	cd . && golangci-lint run --fix
	cd cmd/river && golangci-lint run --fix
	cd riverdriver && golangci-lint run --fix
	cd riverdriver/riverdatabasesql && golangci-lint run --fix
	cd riverdriver/riverpgxv5 && golangci-lint run --fix
	cd rivertype && golangci-lint run --fix

.PHONY: test
test:
	cd . && go test ./... -count=1 -race
	cd cmd/river && go test ./... -count=1 -race
	cd riverdriver && go test ./... -count=1 -race
	cd riverdriver/riverdatabasesql && go test ./... -count=1 -race
	cd riverdriver/riverpgxv5 && go test ./... -count=1 -race
	cd rivertype && go test ./... -count=1 -race

.PHONY: verify
verify:
verify: verify/sqlc

.PHONY: verify/sqlc
verify/sqlc:
	cd riverdriver/riverdatabasesql/internal/dbsqlc && sqlc diff
	cd riverdriver/riverpgxv5/internal/dbsqlc && sqlc diff
