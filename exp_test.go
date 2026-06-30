package math

import (
	"math"
	"testing"
)

func TestExpF64Basic(t *testing.T) {
	input := []float64{0.0, 1.0, 2.0, -1.0, 0.5}
	expected := make([]float64, len(input))
	for i, v := range input {
		expected[i] = math.Exp(v)
	}

	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(output[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], output[i], relErr)
		}
	}
}

func TestExpF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	if math.Abs(output[0]-1.0) > 1e-15 {
		t.Errorf("exp(0) = %f, expected 1.0", output[0])
	}
}

func TestExpF64Inf(t *testing.T) {
	input := []float64{math.Inf(1), math.Inf(-1)}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	if !math.IsInf(output[0], 1) {
		t.Errorf("exp(+Inf) = %f, expected +Inf", output[0])
	}
	if output[1] != 0.0 {
		t.Errorf("exp(-Inf) = %f, expected 0", output[1])
	}
}

func TestExpF64NaN(t *testing.T) {
	input := []float64{math.NaN()}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("exp(NaN) = %f, expected NaN", output[0])
	}
}

func TestExpF64Overflow(t *testing.T) {
	input := []float64{710.0, 750.0, 1000.0}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	for i, v := range output {
		if !math.IsInf(v, 1) {
			t.Errorf("At index %d: expected +Inf for overflow, got %e", i, v)
		}
	}
}

func TestExpF64Underflow(t *testing.T) {
	input := []float64{-750.0, -800.0, -1000.0}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	for i, v := range output {
		if v != 0.0 {
			t.Errorf("At index %d: expected 0 for underflow, got %e", i, v)
		}
	}
}

func TestExpF64Empty(t *testing.T) {
	input := []float64{}
	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	if len(output) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(output))
	}
}

func TestExpF64InPlace(t *testing.T) {
	data := []float64{0.0, 1.0, 2.0, -1.0}
	expected := make([]float64, len(data))
	for i, v := range data {
		expected[i] = math.Exp(v)
	}

	err := ExpF64InPlace(data)
	if err != nil {
		t.Fatalf("ExpF64InPlace failed: %v", err)
	}

	for i := range data {
		absErr := math.Abs(data[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], data[i], relErr)
		}
	}
}

func TestExpF32Basic(t *testing.T) {
	input := []float32{0.0, 1.0, 2.0, -1.0, 0.5}
	expected := make([]float32, len(input))
	for i, v := range input {
		expected[i] = float32(math.Exp(float64(v)))
	}

	output, err := ExpF32(input)
	if err != nil {
		t.Fatalf("ExpF32 failed: %v", err)
	}

	for i := range output {
		relErr := math.Abs(float64((output[i] - expected[i]) / expected[i]))
		if relErr > 1e-6 {
			t.Errorf("At index %d: expected %e, got %e", i, expected[i], output[i])
		}
	}
}

func TestExpF64Large(t *testing.T) {
	n := 10000
	input := make([]float64, n)
	for i := range input {
		input[i] = float64(i%100) / 10.0
	}

	output, err := ExpF64(input)
	if err != nil {
		t.Fatalf("ExpF64 failed: %v", err)
	}

	for i := range output {
		expected := math.Exp(input[i])
		absErr := math.Abs(output[i] - expected)
		relErr := absErr / math.Abs(expected)
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected, output[i], relErr)
		}
	}
}

func BenchmarkExpF64_100(b *testing.B) {
	input := make([]float64, 100)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ExpF64(input)
	}
}

func BenchmarkExpF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ExpF64(input)
	}
}

func BenchmarkExpF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ExpF64(input)
	}
}

func BenchmarkExpF64_100000(b *testing.B) {
	input := make([]float64, 100000)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ExpF64(input)
	}
}

func BenchmarkExpF64Pure_100(b *testing.B) {
	input := make([]float64, 100)
	output := make([]float64, 100)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Exp(input[i])
		}
	}
}

func BenchmarkExpF64Pure_1000(b *testing.B) {
	input := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Exp(input[i])
		}
	}
}

func BenchmarkExpF64Pure_10000(b *testing.B) {
	input := make([]float64, 10000)
	output := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i%100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Exp(input[i])
		}
	}
}
