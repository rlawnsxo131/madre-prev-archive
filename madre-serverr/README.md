# madre-server

### Project Stack

- Golang(workspace)
    - go-chi
    - go-sqlbuilder
- MySQL
- Docker

### Prepare MySQL Database

```shell
$ cd ./docker
$ docker-compose up -d
```

### Start

```shell
# If air is installed
$ air

# If air is not installed
$ go run ./main.go
```
