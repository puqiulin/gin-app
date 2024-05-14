APP_NAME=$(shell basename $$PWD)
APP_VERSION?=$(shell git log -1 --format=%h)
ROOT_PATH=$(shell git rev-parse --show-toplevel)

API_PATH=$(ROOT_PATH)/api
PROTO_FILES=$(shell cd $(API_PATH) && find . -name "*.proto")
PROTOC=cd $(API_PATH) && protoc \
	--proto_path=. \
	--proto_path=../third_party \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	--validate_out=paths=source_relative,lang=go:. \

HEALTH_CMD=curl -f http://127.0.0.1:9999/api/v1/healthz

.PHONY: api
api:
	$(PROTOC) $(PROTO_FILES)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: wire
wire: tidy
	cd pkg/wire && wire

.PHONY: run
run: docker-deps-up
	cd web && pnpm run dev && go run .

.PHONY: run-frontend
run-frontend:
	cd web && pnpm run dev

.PHONY: run-backend
run-backend: wire
	go run .

.PHONY: test
test:
	go test ./...

.PHONY: docker-plugin
docker-plugin:
	docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

.PHONY: docker-deps-up
docker-deps-up:
	docker compose -f docker-compose-deps.yml up -d

.PHONY: docker-deps-down
docker-deps-down:
	docker compose -f docker-compose-deps.yml down

.PHONY: docker-up
docker-up: docker-deps-up
	docker compose -f docker-compose.yml up -d --build

.PHONY: docker-down
docker-down:
	docker compose -f docker-compose.yml down
