.PHONY: start compose router generate

all: download-router generate

start:
	$(MAKE) service & \
	$(MAKE) router & \
	wait

service:
	go run ./cmd/service/main.go

router:
	./router/router -config router.config.yaml

client:
	go run ./cmd/client/main.go

download-router: ./router/router

./router/router:
	npx wgc@latest router download-binary

compose:
	npx wgc@latest router compose -i ./router.compose.yaml -o ./router.execution.config.json

generate-proto:
	npx wgc@latest grpc-service generate -i ./pkg/graph/schema.graphql -o ./pkg/proto/service/v1 --go-package "github.com/wundergraph/service/pkg/generated/service/v1" Service

generate-buf:
	buf generate

generate: generate-proto compose generate-buf