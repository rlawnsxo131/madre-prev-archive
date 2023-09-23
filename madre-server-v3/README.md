# madre-server-v3

### Project Stack

- PostgreSQL
- Go
  - go-chi/chi
  - jackc/pgx
  - rs/zerolog

### Prepare PostgreSQL Database

```shell
$ cd ./docker
$ docker compose up -d
```

### Start

```shell
# If air is installed
$ air

# If air is not installed
$ go run ./main.go
```

### Previous Repository

- v2: <https://github.com/rlawnsxo131/madre-server-v2>
- v1: <https://github.com/rlawnsxo131/madre-server>
