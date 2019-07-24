package main

import (
	"context"
	"testing"

	pb "github.com/jspc/sine-service/service"
)

func TestService_GetGraph(t *testing.T) {
	fullReq := &pb.Sine{
		Frequency:  440.0,
		SampleRate: 19.5,
		Length:     10,
		Multiplier: 5,
	}

	for _, test := range []struct {
		name        string
		redis       RedisClient
		input       *pb.Sine
		expectError bool
	}{
		{"Served from cache", dummyClient{exists: true}, fullReq, false},
		{"Generate data for first time", dummyClient{}, fullReq, false},
		{"Redis error", dummyClient{err: true}, fullReq, true},
		{"Missing data", dummyClient{}, &pb.Sine{}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			s := Service{redis: Redis{test.redis}}

			_, err := s.GetGraph(context.TODO(), test.input)
			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}
		})
	}
}
