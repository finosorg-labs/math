package math

import (
	"math"
	"testing"
)

func TestSqrtF64Basic(t *testing.T) {
	input := []float64{0.0, 1.0, 4.0, 9.0, 16.0, 0.25, 2.0}
	expected := make([]float64, len(input))
	for i, v := range input {
		expected[i] = math.Sqrt(v)
	}

	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(output[i] - expected[i])
		if absErr > 1e-15 {
			t.Errorf("At index %d: expected %e, got %e (abs err %e)", i, expected[i], output[i], absErr)
		}
	}
}

func TestSqrtF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}

	if math.Abs(output[0]-0.0) > 1e-15 {
		t.Errorf("sqrt(0) = %f, expected 0.0", output[0])
	}
}

func TestSqrtF64Negative(t *testing.T) {
	input := []float64{-1.0, -4.0, -0.5}
	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}

	for i, v := range output {
		if !math.IsNaN(v) {
			t.Errorf("sqrt(%f) = %f, expected NaN", input[i], v)
		}
	}
}

func TestSqrtF64Inf(t *testing.T) {
	input := []float64{math.Inf(1), math.Inf(-1)}
	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}

	if !math.IsInf(output[0], 1) {
		t.Errorf("sqrt(+Inf) = %f, expected +Inf", output[0])
	}
	if !math.IsNaN(output[1]) {
		t.Errorf("sqrt(-Inf) = %f, expected NaN", output[1])
	}
}

func TestSqrtF64NaN(t *testing.T) {
	input := []float64{math.NaN()}
	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("sqrt(NaN) = %f, expected NaN", output[0])
	}
}

func TestSqrtF64InPlace(t *testing.T) {
	data := []float64{0.0, 1.0, 4.0, 9.0, 16.0}
	expected := []float64{0.0, 1.0, 2.0, 3.0, 4.0}

	err := SqrtF64InPlace(data)
	if err != nil {
		t.Fatalf("SqrtF64InPlace failed: %v", err)
	}

	for i := range data {
		absErr := math.Abs(data[i] - expected[i])
		if absErr > 1e-15 {
			t.Errorf("At index %d: expected %e, got %e (abs err %e)", i, expected[i], data[i], absErr)
		}
	}
}

func TestSqrtF64Empty(t *testing.T) {
	input := []float64{}
	output, err := SqrtF64(input)
	if err != nil {
		t.Fatalf("SqrtF64 failed: %v", err)
	}
	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func TestSqrtF32Basic(t *testing.T) {
	input := []float32{0.0, 1.0, 4.0, 9.0, 16.0}
	expected := make([]float32, len(input))
	for i, v := range input {
		expected[i] = float32(math.Sqrt(float64(v)))
	}

	output, err := SqrtF32(input)
	if err != nil {
		t.Fatalf("SqrtF32 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(float64(output[i] - expected[i]))
		if absErr > 1e-6 {
			t.Errorf("At index %d: expected %e, got %e (abs err %e)", i, expected[i], output[i], absErr)
		}
	}
}

func BenchmarkSqrtF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i) + 1.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SqrtF64(input)
	}
}

func BenchmarkSqrtF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i) + 1.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SqrtF64(input)
	}
}

func BenchmarkSqrtF64_GoStdlib_1000(b *testing.B) {
	input := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i) + 1.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range input {
			output[j] = math.Sqrt(input[j])
		}
	}
}
