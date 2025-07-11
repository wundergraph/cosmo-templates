# Golang Connect RPC Service

This is a template to create a GraphQL Service using Cosmo Connect.
The Stack is Cosmo Router as the GraphQL-to-gRPC Gateway,
and Connect RPC + Fastify as the backend service.

To run this stack in production, we recommend using [Cosmo Schema Registry](https://cosmo-docs.wundergraph.com/overview) for schema management, breaking change detection, metrics and observability.

## Getting Started

```bash
make
make start
```

Access the GraphQL Playground at [http://localhost:3002](http://localhost:3002).

## Development

Change the GraphQL schema in `./src/graph/schema.graphql` and run:

```bash
make generate
```

Implement the updated proto service in `./pkg/service/service.go` and run:

```bash
make start
```

## Notes

- `router.compose.yaml` is the input to generate the `router.execution.config.json` file, you can add more services/subgraphs
- `router.config.yaml` is the configuration for the Router, you can find the full list of configuration options in the [Router Config Reference](https://cosmo-docs.wundergraph.com/router/configuration)
- `router.execution.config.json` is the output of running `npm run generate:router`, don't modify this file manually
- `buf.gen.yaml` is the configuration for generating the connect rpc files
- `buf.yaml` is the configuration for connect rpc
- `pkg/service/service.go` is the implementation of the connect rpc service, you can add your business logic here
- `pkg/graph/schema.graphql` is the source of truth for the GraphQL Schema, modify it to change the API
- `router/router` is the Cosmo Router binary
- `pkg/generated/service/v1/service.pb.go` is the generated code to encode/decode gRPC messages, don't modify this file manually
- `pkg/generated/service/v1/v1connect/service.connect.go` is the generated code to implement the connect rpc service, don't modify this file manually
- `pkg/proto/service/v1/service.proto` is the proto file that is generated from the GraphQL schema, don't modify this file manually
- `pkg/proto/service/v1/mapping.json` is the mapping file between GraphQL types and proto types, don't modify this file manually
- `pkg/proto/service/v1/service.proto.lock.json` is the lock file for the proto generation, ensuring that e.g. deleted field indentifiers are not reused, don't modify this file manually