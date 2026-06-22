package math

import (
	"math"
	"testing"
)

func TestCumsumF64Basic(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := []float64{1.0, 3.0, 6.0, 10.0, 15.0}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	if len(output) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(output))
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCumsumF64Negative(t *testing.T) {
	input := []float64{-1.0, -2.0, -3.0, -4.0, -5.0}
	expected := []float64{-1.0, -3.0, -6.0, -10.0, -15.0}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCumsumF64Mixed(t *testing.T) {
	input := []float64{1.0, -2.0, 3.0, -4.0, 5.0}
	expected := []float64{1.0, -1.0, 2.0, -2.0, 3.0}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCumsumF64NaN(t *testing.T) {
	input := []float64{1.0, 2.0, math.NaN(), 4.0, 5.0}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	if math.Abs(output[0]-1.0) > 1e-10 {
		t.Errorf("output[0]: expected 1.0, got %f", output[0])
	}
	if math.Abs(output[1]-3.0) > 1e-10 {
		t.Errorf("output[1]: expected 3.0, got %f", output[1])
	}
	if !math.IsNaN(output[2]) {
		t.Errorf("output[2]: expected NaN, got %f", output[2])
	}
	if !math.IsNaN(output[3]) {
		t.Errorf("output[3]: expected NaN, got %f", output[3])
	}
	if !math.IsNaN(output[4]) {
		t.Errorf("output[4]: expected NaN, got %f", output[4])
	}
}

func TestCumsumF64Inf(t *testing.T) {
	input := []float64{1.0, math.Inf(1), 3.0, 4.0, 5.0}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	if math.Abs(output[0]-1.0) > 1e-10 {
		t.Errorf("output[0]: expected 1.0, got %f", output[0])
	}
	if !math.IsInf(output[1], 1) {
		t.Errorf("output[1]: expected +Inf, got %f", output[1])
	}
	if !math.IsInf(output[2], 1) {
		t.Errorf("output[2]: expected +Inf, got %f", output[2])
	}
}

func TestCumsumF64Empty(t *testing.T) {
	input := []float64{}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	if len(output) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(output))
	}
}

func TestCumsumF64InPlace(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := []float64{1.0, 3.0, 6.0, 10.0, 15.0}

	err := CumsumF64InPlace(data)
	if err != nil {
		t.Fatalf("CumsumF64InPlace failed: %v", err)
	}

	for i := range data {
		if math.Abs(data[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], data[i])
		}
	}
}

func TestCumsumF32Basic(t *testing.T) {
	input := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := []float32{1.0, 3.0, 6.0, 10.0, 15.0}

	output, err := CumsumF32(input)
	if err != nil {
		t.Fatalf("CumsumF32 failed: %v", err)
	}

	for i := range output {
		if math.Abs(float64(output[i]-expected[i])) > 1e-6 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCumsumF64Large(t *testing.T) {
	n := 10000
	input := make([]float64, n)
	for i := range input {
		input[i] = 1.0
	}

	output, err := CumsumF64(input)
	if err != nil {
		t.Fatalf("CumsumF64 failed: %v", err)
	}

	for i := range output {
		expected := float64(i + 1)
		if math.Abs(output[i]-expected) > 1e-9 {
			t.Errorf("At index %d: expected %f, got %f", i, expected, output[i])
		}
	}
}

func BenchmarkCumsumF64_100(b *testing.B) {
	input := make([]float64, 100)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CumsumF64(input)
	}
}

func BenchmarkCumsumF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CumsumF64(input)
	}
}

func BenchmarkCumsumF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CumsumF64(input)
	}
}

func BenchmarkCumsumF64_100000(b *testing.B) {
	input := make([]float64, 100000)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CumsumF64(input)
	}
}

func BenchmarkCumsumF64Pure_100(b *testing.B) {
	input := make([]float64, 100)
	output := make([]float64, 100)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0.0
		for i := range input {
			sum += input[i]
			output[i] = sum
		}
	}
}

func BenchmarkCumsumF64Pure_1000(b *testing.B) {
	input := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0.0
		for i := range input {
			sum += input[i]
			output[i] = sum
		}
	}
}

func BenchmarkCumsumF64Pure_10000(b *testing.B) {
	input := make([]float64, 10000)
	output := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i % 100) / 100.0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0.0
		for i := range input {
			sum += input[i]
			output[i] = sum
		}
	}
}
