package math

import (
	"math"
	"testing"
)

func TestLogF64Basic(t *testing.T) {
	input := []float64{1.0, 2.0, 10.0, math.E, 0.5}
	expected := make([]float64, len(input))
	for i, v := range input {
		expected[i] = math.Log(v)
	}

	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(output[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 5e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], output[i], relErr)
		}
	}
}

func TestLogF64One(t *testing.T) {
	input := []float64{1.0}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if math.Abs(output[0]) > 1e-15 {
		t.Errorf("log(1) = %e, expected 0.0", output[0])
	}
}

func TestLogF64E(t *testing.T) {
	input := []float64{math.E}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if math.Abs(output[0]-1.0) > 1e-15 {
		t.Errorf("log(e) = %f, expected 1.0", output[0])
	}
}

func TestLogF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if !math.IsInf(output[0], -1) {
		t.Errorf("log(0) = %f, expected -Inf", output[0])
	}
}

func TestLogF64Negative(t *testing.T) {
	input := []float64{-1.0, -2.0, -10.0}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	for i, v := range output {
		if !math.IsNaN(v) {
			t.Errorf("At index %d: log(negative) = %f, expected NaN", i, v)
		}
	}
}

func TestLogF64Inf(t *testing.T) {
	input := []float64{math.Inf(1)}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if !math.IsInf(output[0], 1) {
		t.Errorf("log(+Inf) = %f, expected +Inf", output[0])
	}
}

func TestLogF64NaN(t *testing.T) {
	input := []float64{math.NaN()}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("log(NaN) = %f, expected NaN", output[0])
	}
}

func TestLogF64Empty(t *testing.T) {
	input := []float64{}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	if len(output) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(output))
	}
}

func TestLogF64InPlace(t *testing.T) {
	data := []float64{1.0, 2.0, 10.0, math.E}
	expected := make([]float64, len(data))
	for i, v := range data {
		expected[i] = math.Log(v)
	}

	err := LogF64InPlace(data)
	if err != nil {
		t.Fatalf("LogF64InPlace failed: %v", err)
	}

	for i := range data {
		absErr := math.Abs(data[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 5e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], data[i], relErr)
		}
	}
}

func TestLogF32Basic(t *testing.T) {
	input := []float32{1.0, 2.0, 10.0, 0.5}
	expected := make([]float32, len(input))
	for i, v := range input {
		expected[i] = float32(math.Log(float64(v)))
	}

	output, err := LogF32(input)
	if err != nil {
		t.Fatalf("LogF32 failed: %v", err)
	}

	for i := range output {
		relErr := math.Abs(float64((output[i] - expected[i]) / expected[i]))
		if relErr > 1e-6 {
			t.Errorf("At index %d: expected %e, got %e", i, expected[i], output[i])
		}
	}
}

func TestLogF64Large(t *testing.T) {
	n := 10000
	input := make([]float64, n)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	for i := range output {
		expected := math.Log(input[i])
		absErr := math.Abs(output[i] - expected)
		relErr := absErr / math.Abs(expected)
		if relErr > 5e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected, output[i], relErr)
		}
	}
}

func TestLogF64SmallValues(t *testing.T) {
	input := []float64{1e-10, 1e-100, 1e-300}
	output, err := LogF64(input)
	if err != nil {
		t.Fatalf("LogF64 failed: %v", err)
	}

	for i := range output {
		expected := math.Log(input[i])
		absErr := math.Abs(output[i] - expected)
		if absErr > 1e-13 {
			t.Errorf("At index %d: expected %e, got %e (abs err %e)", i, expected, output[i], absErr)
		}
	}
}

func BenchmarkLogF64_100(b *testing.B) {
	input := make([]float64, 100)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LogF64(input)
	}
}

func BenchmarkLogF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LogF64(input)
	}
}

func BenchmarkLogF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LogF64(input)
	}
}

func BenchmarkLogF64_100000(b *testing.B) {
	input := make([]float64, 100000)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LogF64(input)
	}
}

func BenchmarkLogF64Pure_100(b *testing.B) {
	input := make([]float64, 100)
	output := make([]float64, 100)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Log(input[i])
		}
	}
}

func BenchmarkLogF64Pure_1000(b *testing.B) {
	input := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Log(input[i])
		}
	}
}

func BenchmarkLogF64Pure_10000(b *testing.B) {
	input := make([]float64, 10000)
	output := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i+1) / 10.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := range input {
			output[i] = math.Log(input[i])
		}
	}
}
