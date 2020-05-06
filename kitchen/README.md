# Engineering Challenge Homework

Create a real-time system that emulates the fulfillment of delivery orders for a kitchen.

For detail, check the [requirement](./doc/requirment).

## Quick Start

Go 1.14 or above.  Using the latest Golang package is preferred.
Turn on go module
```
export GO111MODULE=on
```

Build all binaries
```
# copy this folder to $GOPATH/src first
go install ./...
```
It will install your binary to: `$GOPATH/bin`

Or run below command, it will generate binary in cmd folder.
```
cd cmd/kitchen
go build
```


Run all tests
```
go test ./...
```

Run project
```
$GOPATH/bin/kitchen [-conf config.yaml] orders.json
```

or 
```
cd cmd/kitchen
./kitchen [-conf config.yaml] orders.json
```
