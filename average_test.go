package main

import (
	"testing"
)

func Test(t *testing.T) {
    input := map[string]float64{
        "Math":     90,
        "English":  80,
        "physics":  100,
    }

    expected := 90.0
    result := average(input)

    if result != expected {
        t.Errorf("Expected %.2f, but got %.2f", expected, result)
    }
}