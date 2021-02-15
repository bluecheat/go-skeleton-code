Golang Skeleton code
----

## Install & Build
-  go build
```
go mod vendor
go build ./cmd/skeleton
```
-  app info cli
```
skeleton --version
// Skeleton version 1.0
```

## Configuration 

```yaml
host: "localhost"
port: "9000"
rpcport: "8001"
develop: "true" 
log:
  type: "stdout" # stdout, file
  level: "debug" # info,...
database:
  driver: "memdb" # mysql, memdb
  source: "local:Password@tcp(localhost:3306)/skeleton-code?charset=utf8&parseTime=True&loc=Asia%2FSeoul"
```