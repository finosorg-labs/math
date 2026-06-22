package math

import (
	"math"
	"testing"
)

func TestCummaxF64Basic(t *testing.T) {
	input := []float64{1.0, 3.0, 2.0, 5.0, 4.0}
	expected := []float64{1.0, 3.0, 3.0, 5.0, 5.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
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

func TestCummaxF64Ascending(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expected := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCummaxF64Descending(t *testing.T) {
	input := []float64{5.0, 4.0, 3.0, 2.0, 1.0}
	expected := []float64{5.0, 5.0, 5.0, 5.0, 5.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCummaxF64Negative(t *testing.T) {
	input := []float64{-5.0, -2.0, -8.0, -1.0, -3.0}
	expected := []float64{-5.0, -2.0, -2.0, -1.0, -1.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCummaxF64Mixed(t *testing.T) {
	input := []float64{-5.0, 2.0, -8.0, 10.0, 3.0}
	expected := []float64{-5.0, 2.0, 2.0, 10.0, 10.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	for i := range output {
		if math.Abs(output[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCummaxF64NaN(t *testing.T) {
	input := []float64{1.0, 3.0, math.NaN(), 5.0, 4.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
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

func TestCummaxF64Inf(t *testing.T) {
	input := []float64{1.0, math.Inf(1), 3.0, 5.0, 4.0}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
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

func TestCummaxF64Empty(t *testing.T) {
	input := []float64{}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	if len(output) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(output))
	}
}

func TestCummaxF64InPlace(t *testing.T) {
	data := []float64{1.0, 3.0, 2.0, 5.0, 4.0}
	expected := []float64{1.0, 3.0, 3.0, 5.0, 5.0}

	err := CummaxF64InPlace(data)
	if err != nil {
		t.Fatalf("CummaxF64InPlace failed: %v", err)
	}

	for i := range data {
		if math.Abs(data[i]-expected[i]) > 1e-10 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], data[i])
		}
	}
}

func TestCummaxF32Basic(t *testing.T) {
	input := []float32{1.0, 3.0, 2.0, 5.0, 4.0}
	expected := []float32{1.0, 3.0, 3.0, 5.0, 5.0}

	output, err := CummaxF32(input)
	if err != nil {
		t.Fatalf("CummaxF32 failed: %v", err)
	}

	for i := range output {
		if math.Abs(float64(output[i]-expected[i])) > 1e-6 {
			t.Errorf("At index %d: expected %f, got %f", i, expected[i], output[i])
		}
	}
}

func TestCummaxF64Large(t *testing.T) {
	n := 10000
	input := make([]float64, n)
	for i := range input {
		input[i] = float64(i)
	}

	output, err := CummaxF64(input)
	if err != nil {
		t.Fatalf("CummaxF64 failed: %v", err)
	}

	for i := range output {
		expected := float64(i)
		if math.Abs(output[i]-expected) > 1e-9 {
			t.Errorf("At index %d: expected %f, got %f", i, expected, output[i])
		}
	}
}

func BenchmarkCummaxF64_100(b *testing.B) {
	input := make([]float64, 100)
	for i := range input {
		input[i] = float64(i % 50)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CummaxF64(input)
	}
}

func BenchmarkCummaxF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i % 500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CummaxF64(input)
	}
}

func BenchmarkCummaxF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i % 5000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CummaxF64(input)
	}
}

func BenchmarkCummaxF64_100000(b *testing.B) {
	input := make([]float64, 100000)
	for i := range input {
		input[i] = float64(i % 50000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CummaxF64(input)
	}
}

func BenchmarkCummaxF64Pure_100(b *testing.B) {
	input := make([]float64, 100)
	output := make([]float64, 100)
	for i := range input {
		input[i] = float64(i % 50)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		max := input[0]
		output[0] = max
		for i := 1; i < len(input); i++ {
			if input[i] > max {
				max = input[i]
			}
			output[i] = max
		}
	}
}

func BenchmarkCummaxF64Pure_1000(b *testing.B) {
	input := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range input {
		input[i] = float64(i % 500)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		max := input[0]
		output[0] = max
		for i := 1; i < len(input); i++ {
			if input[i] > max {
				max = input[i]
			}
			output[i] = max
		}
	}
}

func BenchmarkCummaxF64Pure_10000(b *testing.B) {
	input := make([]float64, 10000)
	output := make([]float64, 10000)
	for i := range input {
		input[i] = float64(i % 5000)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		max := input[0]
		output[0] = max
		for i := 1; i < len(input); i++ {
			if input[i] > max {
				max = input[i]
			}
			output[i] = max
		}
	}
}
