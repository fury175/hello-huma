# syntax=docker/dockerfile:1

# ---- build stage ----
FROM golang:1.27 AS build

WORKDIR /src

# Cache dependencies separately so source-only changes don't re-download modules.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/hello-huma ./cmd/

# ---- runtime stage ----
FROM scratch

COPY --from=build /out/hello-huma /hello-huma

EXPOSE 3001

# Run unprivileged; numeric ID works on scratch (no /etc/passwd needed).
USER 65532:65532

ENTRYPOINT ["/hello-huma"]
