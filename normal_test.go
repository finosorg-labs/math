package math

import (
	"math"
	"testing"
)

func TestNormalPdfF64Basic(t *testing.T) {
	input := []float64{0.0, 1.0, -1.0, 2.0, -2.0}
	output, err := NormalPdfF64(input)
	if err != nil {
		t.Fatalf("NormalPdfF64 failed: %v", err)
	}

	invSqrt2Pi := 0.3989422804014327
	for i, x := range input {
		expected := invSqrt2Pi * math.Exp(-0.5*x*x)
		if math.Abs(output[i]-expected) > 1e-12 {
			t.Errorf("Index %d: expected %v, got %v", i, expected, output[i])
		}
	}
}

func TestNormalPdfF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := NormalPdfF64(input)
	if err != nil {
		t.Fatalf("NormalPdfF64 failed: %v", err)
	}

	expected := 0.3989422804014327
	if math.Abs(output[0]-expected) > 1e-15 {
		t.Errorf("pdf(0) should be %v, got %v", expected, output[0])
	}
}

func TestNormalPdfF64Symmetry(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0}
	inputNeg := []float64{-1.0, -2.0, -3.0}

	output, _ := NormalPdfF64(input)
	outputNeg, _ := NormalPdfF64(inputNeg)

	for i := range input {
		if math.Abs(output[i]-outputNeg[i]) > 1e-15 {
			t.Errorf("PDF should be symmetric: pdf(%v) != pdf(%v)", input[i], inputNeg[i])
		}
	}
}

func TestNormalPdfF64InPlace(t *testing.T) {
	data := []float64{0.0, 1.0, -1.0, 2.0}
	invSqrt2Pi := 0.3989422804014327
	expected := make([]float64, len(data))
	for i, x := range data {
		expected[i] = invSqrt2Pi * math.Exp(-0.5*x*x)
	}

	err := NormalPdfF64InPlace(data)
	if err != nil {
		t.Fatalf("NormalPdfF64InPlace failed: %v", err)
	}

	for i := range data {
		if math.Abs(data[i]-expected[i]) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected[i], data[i])
		}
	}
}

func TestNormalPdfF64Empty(t *testing.T) {
	input := []float64{}
	output, err := NormalPdfF64(input)
	if err != nil {
		t.Fatalf("NormalPdfF64 on empty slice failed: %v", err)
	}
	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func TestNormalCdfF64Basic(t *testing.T) {
	input := []float64{0.0, 1.0, -1.0, 2.0, -2.0}
	output, err := NormalCdfF64(input)
	if err != nil {
		t.Fatalf("NormalCdfF64 failed: %v", err)
	}

	invSqrt2 := 0.7071067811865475
	for i, x := range input {
		expected := 0.5 * (1.0 + math.Erf(x*invSqrt2))
		if math.Abs(output[i]-expected) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected, output[i])
		}
	}
}

func TestNormalCdfF64Zero(t *testing.T) {
	input := []float64{0.0}
	output, err := NormalCdfF64(input)
	if err != nil {
		t.Fatalf("NormalCdfF64 failed: %v", err)
	}

	if math.Abs(output[0]-0.5) > 1e-6 {
		t.Errorf("cdf(0) should be 0.5, got %v", output[0])
	}
}

func TestNormalCdfF64Symmetry(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0}
	inputNeg := []float64{-1.0, -2.0, -3.0}

	output, _ := NormalCdfF64(input)
	outputNeg, _ := NormalCdfF64(inputNeg)

	for i := range input {
		sum := output[i] + outputNeg[i]
		if math.Abs(sum-1.0) > 1e-6 {
			t.Errorf("CDF symmetry: cdf(%v) + cdf(%v) should be 1.0, got %v", input[i], inputNeg[i], sum)
		}
	}
}

func TestNormalCdfF64KnownValues(t *testing.T) {
	input := []float64{-3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0}
	expected := []float64{
		0.0013498980316301,
		0.0227501319481792,
		0.1586552539314571,
		0.5,
		0.8413447460685429,
		0.9772498680518208,
		0.9986501019683699,
	}
	output, err := NormalCdfF64(input)
	if err != nil {
		t.Fatalf("NormalCdfF64 failed: %v", err)
	}

	for i := range input {
		if math.Abs(output[i]-expected[i]) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected[i], output[i])
		}
	}
}

func TestNormalCdfF64InPlace(t *testing.T) {
	data := []float64{0.0, 1.0, -1.0, 2.0}
	invSqrt2 := 0.7071067811865475
	expected := make([]float64, len(data))
	for i, x := range data {
		expected[i] = 0.5 * (1.0 + math.Erf(x*invSqrt2))
	}

	err := NormalCdfF64InPlace(data)
	if err != nil {
		t.Fatalf("NormalCdfF64InPlace failed: %v", err)
	}

	for i := range data {
		if math.Abs(data[i]-expected[i]) > 1e-6 {
			t.Errorf("Index %d: expected %v, got %v", i, expected[i], data[i])
		}
	}
}

func TestNormalCdfF64Empty(t *testing.T) {
	input := []float64{}
	output, err := NormalCdfF64(input)
	if err != nil {
		t.Fatalf("NormalCdfF64 on empty slice failed: %v", err)
	}
	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func BenchmarkNormalPdfF64(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	for _, n := range sizes {
		input := make([]float64, n)
		for i := range input {
			input[i] = float64(i%1000)/250.0 - 2.0
		}

		b.Run(string(rune(n)), func(b *testing.B) {
			b.SetBytes(int64(n * 8))
			for i := 0; i < b.N; i++ {
				_, _ = NormalPdfF64(input)
			}
		})
	}
}

func BenchmarkNormalCdfF64(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	for _, n := range sizes {
		input := make([]float64, n)
		for i := range input {
			input[i] = float64(i%1000)/250.0 - 2.0
		}

		b.Run(string(rune(n)), func(b *testing.B) {
			b.SetBytes(int64(n * 8))
			for i := 0; i < b.N; i++ {
				_, _ = NormalCdfF64(input)
			}
		})
	}
}
