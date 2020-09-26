# Backend

## setup
```
$ make
$ make migrate-up
```

## Commands
### start containers
```
$ make
```

### start app server
```
$ make start
```

### go generate
```
$ make generate
```

### migration
```
# create table
$ FILENAME=<filename> make migrate-create

# up
$ make migrate-up

# down
$ make migrate-down
```