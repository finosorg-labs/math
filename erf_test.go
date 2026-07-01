package math

import (
	"math"
	"testing"
)

func TestErfF64Basic(t *testing.T) {
	input := []float64{0.0, 1.0, 2.0, -1.0, 0.5}
	output, err := ErfF64(input)
	if err != nil {
		t.Fatalf("ErfF64 failed: %v", err)
	}

	for i, x := range input {
		expected := math.Erf(x)
		if math.Abs(output[i]-expected) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected, output[i])
		}
	}
}

func TestErfF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := ErfF64(input)
	if err != nil {
		t.Fatalf("ErfF64 failed: %v", err)
	}

	if math.Abs(output[0]) > 1e-15 {
		t.Errorf("erf(0) should be 0, got %v", output[0])
	}
}

func TestErfF64Symmetry(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0}
	inputNeg := []float64{-1.0, -2.0, -3.0}

	output, _ := ErfF64(input)
	outputNeg, _ := ErfF64(inputNeg)

	for i := range input {
		if math.Abs(output[i]+outputNeg[i]) > 1e-15 {
			t.Errorf("erf should be antisymmetric: erf(%v) + erf(%v) = %v", input[i], inputNeg[i], output[i]+outputNeg[i])
		}
	}
}

func TestErfF64InPlace(t *testing.T) {
	data := []float64{0.0, 1.0, -1.0, 2.0}
	expected := make([]float64, len(data))
	for i, x := range data {
		expected[i] = math.Erf(x)
	}

	err := ErfF64InPlace(data)
	if err != nil {
		t.Fatalf("ErfF64InPlace failed: %v", err)
	}

	for i := range data {
		if math.Abs(data[i]-expected[i]) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected[i], data[i])
		}
	}
}

func TestErfF64Empty(t *testing.T) {
	input := []float64{}
	output, err := ErfF64(input)
	if err != nil {
		t.Fatalf("ErfF64 on empty slice failed: %v", err)
	}
	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func BenchmarkErfF64(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	for _, n := range sizes {
		input := make([]float64, n)
		for i := range input {
			input[i] = float64(i%1000)/500.0 - 1.0
		}

		b.Run(string(rune(n)), func(b *testing.B) {
			b.SetBytes(int64(n * 8))
			for i := 0; i < b.N; i++ {
				_, _ = ErfF64(input)
			}
		})
	}
}
