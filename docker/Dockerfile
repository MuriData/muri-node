# ── Build stage ──
FROM golang:1.24-bookworm AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /murid ./cmd/murid/

# ── Final stage ──
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN groupadd -r murid && useradd -r -g murid -d /app -s /sbin/nologin murid
WORKDIR /app

COPY --from=builder /murid /usr/local/bin/murid
COPY docker/docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# Default mount points
RUN mkdir -p /app/keys /app/data && chown -R murid:murid /app

USER murid

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["run"]
