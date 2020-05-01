# Simple webserver

## Prerequisites:  
* Have [Go](http://golang.org) installed
* Have the Go workspace bin folder accessible (`$GOPATH/bin`)

## Installation

``` bash
go get github.com/chrisvdg/simpleserver
```

## Use

```bash
# List all flags
simpleserver --help

# Serves default directory (./public) on a plain http server on the default port (8080).
# If the default directory does not exists, making a request to the server will return 404
simpleserver

# Serves the static subdirectory on port 6060
simpleserver --port 6060 --dir ./static 
```

### TLS

Simpleserver supports TLS if tls-cert and tls-key flags are provided:
```bash
simpleserver --port 6060 --dir ./static --tls-cert certificate_file.cert --tls-key private_key_file.key
```