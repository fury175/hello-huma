# hello-huma

[![Audit (CI)](https://github.com/fury175/hello-huma/actions/workflows/audit.yml/badge.svg?branch=main)](https://github.com/fury175/hello-huma/actions/workflows/audit.yml)

A minimal HTTP API starter template built with [Huma v2](https://huma.rocks), [Chi](https://github.com/go-chi/chi) and [Zap](https://go.uber.org/zap) — a small, opinionated starting point for a Go microservice.

## Stack

| Concern | Library |
| --- | --- |
| HTTP framework | [Huma v2](https://github.com/danielgtaylor/huma) (chi adapter) |
| Router | [go-chi/chi v5](https://github.com/go-chi/chi) |
| Logging | [uber-go/zap](https://go.uber.org/zap) |
| Config | [caarlos0/env v11](https://github.com/caarlos0/env) + godotenv autoload |
| Dev loop | [air](https://github.com/air-verse/air) (live reload, config in `.air.toml`) |

## Requirements

- Go 1.26+ (minimum supported; see the `go` directive in `go.mod`)
- Docker (optional, for containerized runs)
- air (optional, for live reload): `go install github.com/air-verse/air@latest`

## Getting started

```bash
cp .env.example .env   # ADDR=:3001, ENV=development
```

Run with live reload:

```bash
air
```

...or build and run directly:

```bash
go build -o bin/main ./cmd/ && ./bin/main
```

The server listens on `:3001` by default.

## Endpoints

| Method | Path             | Description            |
| ------ | ---------------- | ---------------------- |
| GET    | `/api/v1/health` | Liveness / health probe |

API docs (Scalar): <http://localhost:3001/docs>
OpenAPI spec: <http://localhost:3001/openapi.json>

### Example

```bash
curl http://localhost:3001/api/v1/health
# {"status":"ok","environment":"development","success":true}
```

## Configuration

Environment variables (also read from `.env` automatically):

| Variable | Default       | Description                          |
| -------- | ------------- | ------------------------------------ |
| `ADDR`   | `:3001`       | Address the server listens on        |
| `ENV`    | `development` | Environment name, reported by health |

## Tests

```bash
go test ./...
```

## Docker

```bash
docker build -t hello-huma .
docker run --rm -p 3001:3001 hello-huma
```

The runtime stage is built on `scratch` with a static (`CGO_ENABLED=0`) binary and runs unprivileged — the image is ~12 MB.

Note: `scratch` ships without a CA store, so the Dockerfile copies the CA bundle from the builder stage (`/etc/ssl/certs/ca-certificates.crt`). Without it, any outbound HTTPS call from the container fails with a certificate-verification error.

Override config at runtime:

```bash
docker run --rm -p 3001:3001 -e ADDR=:8080 -e ENV=production hello-huma
```

## Releases

Versions are tagged and released automatically by [release-please](https://github.com/googleapis/release-please) — see [CHANGELOG.md](CHANGELOG.md).

## Project layout

```
.
├── cmd/
│   ├── main.go        # entrypoint: config load + run
│   ├── api.go         # chi + huma wiring, graceful shutdown
│   ├── health.go      # GET /api/v1/health
│   ├── health_test.go # unit tests
│   └── types.go       # application + config types
├── internal/
│   ├── apperror/      # huma error helpers + middleware
│   └── logger/        # zap logger singleton
├── .github/
│   └── workflows/     # CI (audit), release-please, auto version bump
├── .air.toml          # air (live reload) config
├── .env.example
├── .gitignore
├── .dockerignore
├── CHANGELOG.md       # maintained by release-please
├── Dockerfile
├── LICENSE            # MIT
├── go.mod
└── go.sum
```
