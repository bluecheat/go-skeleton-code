IOT Broker API
----
> Collects IOT data and provides multiple streams
> 1. Save iot row-data as CSV file
> 2. Send iot row-data to the Backoffice real-time dashboard
> 3. Change the iot data group to hash and send it to the ReapChain

## Install & Build
-  go build
```
go mod vendor
go build cmd/iotbroker
```
-  proto build (optional)
```
mkdir -p $GOPATH/google/api 
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > $GOPATH/google/api/annotations.proto     
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > $GOPATH/google/api/http.proto
```

## Configuration 

```yaml
port: "11000" // restful port
rpcport: "11001" // grpc port
develop: "true" // debug mode
stream:
  blockchainEvent:  // blockchain stream option
    brokers:
      - "localhost:9092"
      - "localhost:9092"
      - "localhost:9092"
    topic: "reapchain.fabric.fct.transaction"
    maxsize: 1000000 # byte
    timeout: 10 # sec
  backofficeEvent: // backoffice stream option
    brokers:
      - "localhost:9092"
      - "localhost:9092"
      - "localhost:9092"
    topic: "reapchain.iot.fct.device"
  file:
    path: "./"
log: // log config
  type: "stdout"  // stdout, file
  level: "debug"  // trace, debug, info, error
```