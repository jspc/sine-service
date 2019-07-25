package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/jspc/sine-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = flag.String("address", "sine.ori.jspc.pw:443", "Location of sine service")

	frequency  = flag.Float64("frequency", 445.0, "Sine wave frequency")
	sampleRate = flag.Float64("sample", 88.8, "Sine wave sample rate")
	multiplier = flag.Float64("multiplier", 10, "Point multiplier")
	length     = flag.Int64("length", 250, "Number of points to plot")
	outputFile = flag.String("file", "graph.png", "File to write graph to")

	verbose = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	realMain(*addr, *frequency, *sampleRate, *multiplier, *length, *outputFile, *verbose)
}

// realMain performs the brunt of the work; we split it to make testing simpler
func realMain(addr string, freq, sample, multiplier float64, length int64, output string, verbose bool) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewSineServiceClient(conn)

	sg, err := client.GetGraph(context.Background(), &pb.Sine{
		Frequency:  freq,
		SampleRate: sample,
		Multiplier: multiplier,
		Length:     length,
	})

	if err != nil {
		panic(err)
	}

	if verbose {
		log.Printf("%#v", sg)
	}

	decoded, err := base64.StdEncoding.DecodeString(sg.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(output, decoded, 0640)
	if err != nil {
		panic(err)
	}
}
