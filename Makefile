.PHONY: install
install:
	make install_default_tools
	make install_linters
	make install_spanner_tools
	make install_wire
	make install_mock
	make install_stringer

.PHONY: install_default_tools
install_default_tools:
	go install golang.org/x/tools/cmd/goimports@v0.25.0
	go install github.com/daixiang0/gci@v0.13.5
	go install github.com/mikefarah/yq/v4@v4.44.5

.PHONY: install_linters
install_linters:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1

.PHONY: install_spanner_tools
install_spanner_tools:
	go install github.com/cloudspannerecosystem/spanner-cli@v0.10.9
	go install cuelang.org/go/cmd/cue@v0.6.0
	go install -tags 'spanner,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1
	go install github.com/daichirata/hammer@v0.6.5

.PHONY: install_wire
install_wire:
	go install github.com/google/wire/cmd/wire@v0.6.0

.PHONY: install_mock
install_mock:
	go install go.uber.org/mock/mockgen@v0.4.0

.PHONY: install_stringer
install_stringer:
	go install golang.org/x/tools/cmd/stringer@v0.25.0

.PHONY: gqlgen
gqlgen:
	go run -x github.com/99designs/gqlgen generate --verbose

.PHONY: entgen
entgen:
	go run -x -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target=./pkg/adapter/entgen ./schema/db/ent

.PHONY: docker_up
docker_up:
	docker network create sns-server || true
	docker compose up -d --build
	
.PHONY: docker_down
docker_down:
	docker compose down

.PHONY: migrate
migrate:
	go run cmd/migration/main.go

.PHONY: run_api
run_api:
	go run ./main.go 
