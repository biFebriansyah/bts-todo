# bts-todoapp

### Prerequisites

-  Go (v1.23+)
-  Postgres instance (local or cloud)

### Project setup

```bash
# Clone the repository
$ git clone https://github.com/biFebriansyah/bts-todo.git

# Install Package
$ go get -u ./...

# using Make
$ make install
```

### Compile and run the project

```bash
# development
$ go run *.go

# watch mode
$ make run

# production mode
$ make build
```

### Run Migration

```bash
# create migration file
$ make migrate-init name=todoapp-name-migration

# exec migration
$ make migrate-up

# reset migration
$ make migrate-reset
```

## Authors

-  [@biFebriansyah](https://www.github.com/biFebriansyah)
