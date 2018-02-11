# Simple webserver

## Prerequisites:  
* Have [Go](http://golang.org) installed
* Have the Go workspace bin folder accessible (`$GOPATH/bin`)

## Installation

``` bash
go get github.com/scarletdraped/simpleserver
```

## Use

```bash
# list all flags
simpleserver --help

# Serves current directory on a plain http server
# on the default port (8080)
simpleserver

# Serves the static subdirectory on port 6060
simpleserver --port 6060 --dir ./static 
```

### TLS

Simpleserver supports TLS if tls-cert and tls-key flags are provided:
```bash
simpleserver --port 6060 --dir ./static --tls-cert certificate_file.cert --tls-key private_key_file.key
```