package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSine_Intialise(t *testing.T) {
	for _, test := range []struct {
		name        string
		sine        *Sine
		expectError bool
		expectFPS   float64
	}{
		{"Happy path", &Sine{Frequency: 1.0, SampleRate: 100.0, Multiplier: 8.0, Length: 5}, false, 0.06283},
		{"Missing Freq", &Sine{SampleRate: 100.0, Multiplier: 8.0, Length: 5}, true, 0.0},
		{"Missing SR", &Sine{Frequency: 100.0, Multiplier: 8.0, Length: 5}, true, 0.0},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("unexpected error: %+v", err)
				}
			}()

			err := test.sine.Initialise()
			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if fmt.Sprintf("%.5f", test.expectFPS) != fmt.Sprintf("%.5f", test.sine.freqPerSample) {
				t.Errorf("expected %.5f, received %.5f", test.expectFPS, test.sine.freqPerSample)
			}
		})
	}
}

func TestSine_Graph(t *testing.T) {
	s := &Sine{
		Frequency:  1000.0,
		SampleRate: 44.1,
		Multiplier: 15.0,
		Length:     10,
	}

	s.Initialise()

	g := s.Graph()
	expect := []float64{
		-13.396487277725198, 12.053084119853981, 2.5520895047483227, -14.349249438748782, 10.35821354384183, 5.029760188537628, -14.88358809900192, 8.361297903297471, 7.360763280052259, -14.983921965974695,
	}

	if !reflect.DeepEqual(expect, g) {
		t.Errorf("expected:\n\t%+v\nreceived:\n\t%+v", expect, g)
	}
}

func TestSine_RequestID(t *testing.T) {
	s := &Sine{
		Frequency:  1000.0,
		SampleRate: 44.1,
		Multiplier: 15.0,
		Length:     10,
	}

	s.Initialise()

	reqID := s.RequestID()
	expect := "H4sIAAAAAAAA/zI0MDDQM0ACNTUmJnqGKAKGpmgqDA0AAQAA//8ySF7ONwAAAA=="

	if expect != reqID {
		t.Errorf("expected %q, received %q", expect, reqID)
	}
}

var g []float64

func benchmarkSine(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := &Sine{
			Frequency:  1000.0,
			SampleRate: 44.1,
			Multiplier: 15.0,
			Length:     i,
		}

		s.Initialise()
		g = s.Graph()
	}
}

func BenchmarkSine_Graph1(b *testing.B)        { benchmarkSine(1, b) }
func BenchmarkSine_Graph10(b *testing.B)       { benchmarkSine(10, b) }
func BenchmarkSine_Graph100(b *testing.B)      { benchmarkSine(100, b) }
func BenchmarkSine_Graph1000(b *testing.B)     { benchmarkSine(1000, b) }
func BenchmarkSine_Graph10000(b *testing.B)    { benchmarkSine(10000, b) }
func BenchmarkSine_Graph100000(b *testing.B)   { benchmarkSine(100000, b) }
func BenchmarkSine_Graph1000000(b *testing.B)  { benchmarkSine(1000000, b) }
func BenchmarkSine_Graph10000000(b *testing.B) { benchmarkSine(10000000, b) }
