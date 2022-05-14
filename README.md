# madre-server-v2

### Project Stack

- PostgreSQL
- Go
  - gorilla/mux
  - jmoiron/sqlx
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

### V1 Repository

<https://github.com/rlawnsxo131/madre-server>
