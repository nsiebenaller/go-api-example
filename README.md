# go-api-example

Example Go REST API using the Gin Web Framework

### Commands

- Install and Build

```bash
go install
```

- Start dev server

```bash
# Make sure `air` is installed
go install github.com/air-verse/air@latest
# Start dev server
air
```

- Docker scripts

```bash
# Startup PostgreSQL docker database
docker compose -f docker-compose.yml up --force-recreate --build -d
# Shutdown PostgreSQL docker database
docker compose -f docker-compose.yml down --remove-orphans --volumes --rmi all
# Stop PostgreSQL docker database
docker compose -f docker-compose.yml stop
```

### Project Structure

| Directory   | Description                                               |
| ----------- | --------------------------------------------------------- |
| controllers | Route handlers                                            |
| docker      | Docker configuration files for local development          |
| models      | Structs for database access                               |
| router      | HTTP router for connecting routes with controllers        |
| services    | Shared functionality across controllers (database access) |
| tests       | HTTP requests for testing purposes                        |
