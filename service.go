package main

import (
	"context"
	"log"

	pb "github.com/jspc/sine-service/service"
)

// Service implements the SineServiceServer interface
// defined in ./service, which is generated from
// ./grpc/sine-service.proto
type Service struct {
	redis Redis
}

// GetGraph takes some sine input and returns a base64 encoded
// png with that input plotted on a graph.
//
// The supplied context is ignored for now; later implementations
// will expand this
func (s Service) GetGraph(_ context.Context, input *pb.Sine) (sg *pb.SineGraph, err error) {
	// Todo: consider reflecting these values in
	sine := Sine{
		Frequency:  input.Frequency,
		SampleRate: input.SampleRate,
		Multiplier: input.Multiplier,
		Length:     int(input.Length),
	}

	err = sine.Initialise()
	if err != nil {
		return
	}

	exists, err := s.redis.Exists(sine)
	if err != nil {
		return
	}

	var img string
	if exists {
		img, err = s.redis.Read(sine)

		sg = &pb.SineGraph{
			Body:    img,
			Message: "from cache",
		}

		return
	}

	graph := sine.Graph()
	chart, err := NewChart(graph)
	if err != nil {
		return
	}

	sg = &pb.SineGraph{
		Body: chart.Base64(),
	}

	go func() {
		ok, err := s.redis.Write(sine, chart)
		if err != nil || !ok {
			log.Printf("Error writing %q to redis, ok: %v, error: %+v",
				sine.RequestID(), ok, err)
		}
	}()

	return
}
