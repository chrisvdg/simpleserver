# Simple webserver

## Prerequisites:  
* Have [Go](http://golang.org) installed
* Have the Go bin folder accessible

## Installation

``` bash
go get github.com/scarletdraped/simpleserver
```

## Use

```bash
# Serves current directory on port 8080
simpleserver 

# Serves static subdirectory on port 6060
simpleserver --port 6060 --dir ./static 
```