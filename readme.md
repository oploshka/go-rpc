cd ./project_root
go mod download
go mod vendor
go mod verify

cd ..

docker-compose up -d

make migr-status
make migr-up

# Go Rpc

## Linter

https://golangci-lint.run

## GoFmt

пройтись по всему проекту

```shell
gofmt -s -w .
```


```shell
go build main.go && main.exe
```
