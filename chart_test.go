package main

import (
	"testing"
)

func TestNewChart(t *testing.T) {
	for _, test := range []struct {
		name        string
		points      []float64
		expectError bool
	}{
		{"happy path", []float64{0.0, 0.1, 0.1, 0.2, 0.1, 0.1, 0.0}, false},
		{"empty points", []float64{}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewChart(test.points)
			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}
		})
	}
}

func TestChart_Base64(t *testing.T) {
	// Catch errors- this is a purely happy path test

	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	c, _ := NewChart([]float64{0.1, 0.2, 0.1})

	c.Base64()
}
