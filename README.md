## Go Rpc

```shell
#install vendor & docker-compose up -d
make init

#database migrations (status)
make migr-status

#database migrations (process)
make migr-up
```

+ [localhost:8080/api-form](http://localhost:8080/api-form) - форма для отладки

## Linter
[https://golangci-lint.run](https://golangci-lint.run)

```shell
gofmt -s -w .
```