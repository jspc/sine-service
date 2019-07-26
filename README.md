[![codecov](https://codecov.io/gh/jspc/sine-service/branch/master/graph/badge.svg)](https://codecov.io/gh/jspc/sine-service)

# Sine Service

The sine service is a gRPC driven service which, given a set of inputs such as frequency and sample rate, will return a graph of the resultant sine wave.

Graphs are plotted and drawn server-side; this is returned to clients in a protobuf which looks like:

```protobuf
// SineGraph is the an encoded file with optional debug text
message SineGraph {
  // Body contains a base64 encoded png
  string Body = 1;

  // Message is optional debug
  string Message = 2;
}
```

## Runbook

For incident management, please see [doc/RUNBOOK.md](doc/RUNBOOK.md)

## Client

A client is provided for this service, which may be installed by:

```bash
$ go get github.com/jspc/sine-service/sine-cli
```

This client can then be invoked as:

```bash
$ sine-cli -h
Usage of sine-cli:
  -address string
        Location of sine service (default "sine.ori.jspc.pw:443")
  -file string
        File to write graph to (default "graph.png")
  -frequency float
        Sine wave frequency (default 445)
  -length int
        Number of points to plot (default 250)
  -multiplier float
        Point multiplier (default 10)
  -sample float
        Sine wave sample rate (default 88.8)
  -v    Verbose output

```

The default values will write the file `graph.png` which looks like:

![sine wave plotted on graph](doc/sample-graph.png?raw=true "Sine wave plotted on graph")

## Architecture

![sine-service architecture diagram](doc/sine-servicedrawio.png?raw=true "Sine Service architecture diagram")

The sine service uses redis to cache data, to avoid redrawing potentially complexe graphs multiple times.

## Testing

This project strives for 100% code coverage, but recognises that this is not always possible. Tests, with coverage, may be run with

```bash
$ go test -covermode=count -coverprofile=count.out -v
```
