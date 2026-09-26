# AMA Rooms API (Go)

Backend for "Ask Me Anything" rooms, started in Go during a Rocketseat Go + React event. People create rooms, post questions, react to them and mark them as answered, with live updates over WebSockets.

> **Status: work in progress.** The database schema, migrations, generated queries and the full route table are done. The HTTP handlers are still empty stubs.

## Planned endpoints

| Method | Route | What it does |
| --- | --- | --- |
| `GET` | `/subscribe/{room_id}` | WebSocket: live updates for a room |
| `POST` | `/api/rooms` | Create a room |
| `GET` | `/api/rooms` | List rooms |
| `POST` | `/api/rooms/{room_id}/messages` | Post a question |
| `GET` | `/api/rooms/{room_id}/messages` | List a room's questions |
| `GET` | `/api/rooms/{room_id}/messages/{message_id}` | Get one question |
| `PATCH` / `DELETE` | `/api/rooms/{room_id}/messages/{message_id}/react` | Add or remove a reaction |
| `PATCH` | `/api/rooms/{room_id}/messages/{message_id}/answer` | Mark a question as answered |

## Stack

- **Go** with [chi](https://github.com/go-chi/chi) for routing and [gorilla/websocket](https://github.com/gorilla/websocket)
- **PostgreSQL** via [pgx](https://github.com/jackc/pgx), running in Docker Compose
- [sqlc](https://sqlc.dev) to generate type-safe queries from SQL
- [tern](https://github.com/jackc/tern) for migrations

## Running

Requires Go 1.22+, Docker, `sqlc` and `tern`.

```bash
# Start Postgres (reads WSRS_DATABASE_* from .env)
docker compose up -d

# Run migrations and generate queries
go generate ./...

# Start the API
go run ./cmd/wsrs
```
