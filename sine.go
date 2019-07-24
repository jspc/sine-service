package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"math"
)

const (
	phase = 0.0
)

type Sine struct {
	Frequency  float64 `json:"frequency"`
	SampleRate float64 `json:"sample_rate"`
	Multiplier float64 `json:"multiplier"`
	Length     int     `json:"length"`

	freqPerSample float64
	phase         float64
}

// Initialise acts on an instance of Sine, validates the
// fields are as expected, and generates some mathematic
// values to operate upon
func (s *Sine) Initialise() error {
	if s.Frequency == 0 {
		return fmt.Errorf("Frequency must not be 0")
	}

	if s.SampleRate == 0 {
		return fmt.Errorf("SampleRate must not be 0")
	}

	if s.Multiplier == 0 {
		s.Multiplier = 1
	}

	s.freqPerSample = 2 * math.Pi * s.Frequency / s.SampleRate
	s.phase = phase

	return nil
}

// Next returns the next value in the sine wave
func (s *Sine) Next() float64 {
	s.phase = s.phase + s.freqPerSample

	return float64(math.Sin(s.phase) * s.Multiplier)
}

// Graph will return a slice of floats representing the values
// in our sine wave
func (s *Sine) Graph() (g []float64) {
	g = make([]float64, s.Length)

	for i := 0; i < s.Length; i++ {
		g[i] = s.Next()
	}

	return
}

func (s Sine) RequestID() string {
	str := fmt.Sprintf("%.12f||%.12f||%.12f||%d",
		s.Frequency,
		s.SampleRate,
		s.Multiplier,
		s.Length,
	)

	b := bytes.Buffer{}
	w := gzip.NewWriter(&b)
	w.Write([]byte(str))
	w.Close()

	return base64.StdEncoding.EncodeToString(b.Bytes())
}
