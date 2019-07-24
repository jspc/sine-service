# Sine Service

The sine service is a gRPC driven service which, given a set of inputs such as frequency and sample rate, will return a graph of the resultant sine wave.

## Architecture

![sine-service architecture diagram](doc/sine-servicedrawio.png?raw=true "Sine Service architecture diagram")

The sine service uses redis to cache data, to avoid redrawing potentially complexe graphs multiple times.
