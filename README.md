## Migoro

![go-release](https://github.com/brownhounds/migoro/actions/workflows/go.yml/badge.svg)
![docker-build](https://github.com/brownhounds/migoro/actions/workflows/docker.yml/badge.svg)

Database Migrator for Postgres and SQLite

## Install

### Build From Source

Dependencies: go >= 1.22.2

1. Clone repository
2. Run: `go get`
3. Build:

```bash
go generate ./...
go build -ldflags="-s -w" -o ./bin/migoro main.go
```

### Install Script LINUX/amd64

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/brownhounds/migoro/v0.1.3/tools/install-linux-amd64.sh)"
```

### Install Script LINUX/arm64

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/brownhounds/migoro/v0.1.3/tools/install-linux-arm64.sh)"
```

### Post Install/LINUX

Add following to `.bashrc` or `.zshrc`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Source `profile` file or restart your terminal session.

## Other Platforms

Refer to latest release page: [Release](https://github.com/brownhounds/migoro/releases)

## Docker

Docker image available on: [Docker Hub](https://hub.docker.com/r/brownhounds/migoro)

Example `docker-compose.yml`:

```yml
services:
  postgres:
    image: postgres:alpine
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U admin"]
      interval: 5s
      retries: 5
      timeout: 10s
    ports:
      - 5432:5432

  migrator:
    image: brownhounds/migoro:latest
    environment:
      - SQL_DRIVER=postgres
      - SQL_HOST=postgres
      - SQL_PORT=5432
      - SQL_USER=admin
      - SQL_PASSWORD=admin
      - SQL_DB=test
      - SQL_DB_SCHEMA=public
      - SQL_SSL=disable
      - MIGRATION_DIR=/go/bin/migrations
      - MIGRATION_TABLE=migrations
      - MIGRATION_SCHEMA=platform
    volumes:
      - ./migrations:/go/bin/migrations
    depends_on:
      postgres:
        condition: service_healthy
```
