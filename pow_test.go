package math

import (
	"math"
	"testing"
)

func TestPowF64Basic(t *testing.T) {
	base := []float64{2.0, 3.0, 10.0, 0.5, 2.0}
	exponent := []float64{3.0, 2.0, 1.0, -1.0, 0.5}
	expected := make([]float64, len(base))
	for i := range base {
		expected[i] = math.Pow(base[i], exponent[i])
	}

	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(output[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], output[i], relErr)
		}
	}
}

func TestPowF64ZeroExponent(t *testing.T) {
	base := []float64{2.0, -5.0, 0.0, math.NaN()}
	exponent := []float64{0.0, 0.0, 0.0, 0.0}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	for i, v := range output {
		if v != 1.0 {
			t.Errorf("pow(%f, 0) = %f, expected 1.0", base[i], v)
		}
	}
}

func TestPowF64BaseOne(t *testing.T) {
	base := []float64{1.0, 1.0, 1.0}
	exponent := []float64{5.0, -3.0, math.NaN()}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	for i, v := range output {
		if v != 1.0 {
			t.Errorf("pow(1, %f) = %f, expected 1.0", exponent[i], v)
		}
	}
}

func TestPowF64ZeroBase(t *testing.T) {
	base := []float64{0.0, 0.0, 0.0}
	exponent := []float64{2.0, -2.0, 0.0}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	if output[0] != 0.0 {
		t.Errorf("pow(0, 2) = %f, expected 0.0", output[0])
	}
	if !math.IsInf(output[1], 1) {
		t.Errorf("pow(0, -2) = %f, expected +Inf", output[1])
	}
	if output[2] != 1.0 {
		t.Errorf("pow(0, 0) = %f, expected 1.0", output[2])
	}
}

func TestPowF64NegativeBase(t *testing.T) {
	base := []float64{-2.0}
	exponent := []float64{0.5}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("pow(-2, 0.5) = %f, expected NaN", output[0])
	}
}

func TestPowF64NaN(t *testing.T) {
	base := []float64{math.NaN(), 2.0}
	exponent := []float64{2.0, math.NaN()}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("pow(NaN, 2) = %f, expected NaN", output[0])
	}
	if !math.IsNaN(output[1]) {
		t.Errorf("pow(2, NaN) = %f, expected NaN", output[1])
	}
}

func TestPowF64MismatchedLength(t *testing.T) {
	base := []float64{2.0, 3.0}
	exponent := []float64{2.0}
	_, err := PowF64(base, exponent)
	if err == nil {
		t.Fatal("Expected error for mismatched lengths, got nil")
	}
}

func TestPowF64InPlace(t *testing.T) {
	base := []float64{2.0, 3.0, 4.0}
	exponent := []float64{2.0, 2.0, 2.0}
	expected := []float64{4.0, 9.0, 16.0}

	err := PowF64InPlace(base, exponent)
	if err != nil {
		t.Fatalf("PowF64InPlace failed: %v", err)
	}

	for i := range base {
		absErr := math.Abs(base[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], base[i], relErr)
		}
	}
}

func TestPowScalarF64Basic(t *testing.T) {
	base := []float64{2.0, 3.0, 4.0, 5.0}
	exponent := 2.0
	expected := []float64{4.0, 9.0, 16.0, 25.0}

	output, err := PowScalarF64(base, exponent)
	if err != nil {
		t.Fatalf("PowScalarF64 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(output[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], output[i], relErr)
		}
	}
}

func TestPowScalarF64ZeroExponent(t *testing.T) {
	base := []float64{2.0, -5.0, 0.0}
	exponent := 0.0
	output, err := PowScalarF64(base, exponent)
	if err != nil {
		t.Fatalf("PowScalarF64 failed: %v", err)
	}

	for i, v := range output {
		if v != 1.0 {
			t.Errorf("pow(%f, 0) = %f, expected 1.0", base[i], v)
		}
	}
}

func TestPowScalarF64InPlace(t *testing.T) {
	base := []float64{2.0, 3.0, 4.0}
	exponent := 3.0
	expected := []float64{8.0, 27.0, 64.0}

	err := PowScalarF64InPlace(base, exponent)
	if err != nil {
		t.Fatalf("PowScalarF64InPlace failed: %v", err)
	}

	for i := range base {
		absErr := math.Abs(base[i] - expected[i])
		relErr := absErr / math.Abs(expected[i])
		if relErr > 1e-4 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], base[i], relErr)
		}
	}
}

func TestPowF64Empty(t *testing.T) {
	base := []float64{}
	exponent := []float64{}
	output, err := PowF64(base, exponent)
	if err != nil {
		t.Fatalf("PowF64 failed: %v", err)
	}
	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func TestPowF32Basic(t *testing.T) {
	base := []float32{2.0, 3.0, 10.0}
	exponent := []float32{3.0, 2.0, 1.0}
	expected := make([]float32, len(base))
	for i := range base {
		expected[i] = float32(math.Pow(float64(base[i]), float64(exponent[i])))
	}

	output, err := PowF32(base, exponent)
	if err != nil {
		t.Fatalf("PowF32 failed: %v", err)
	}

	for i := range output {
		absErr := math.Abs(float64(output[i] - expected[i]))
		relErr := absErr / math.Abs(float64(expected[i]))
		if relErr > 1e-3 {
			t.Errorf("At index %d: expected %e, got %e (rel err %e)", i, expected[i], output[i], relErr)
		}
	}
}

func BenchmarkPowF64_1000(b *testing.B) {
	base := make([]float64, 1000)
	exponent := make([]float64, 1000)
	for i := range base {
		base[i] = float64(i%100) + 1.0
		exponent[i] = 2.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = PowF64(base, exponent)
	}
}

func BenchmarkPowScalarF64_1000(b *testing.B) {
	base := make([]float64, 1000)
	for i := range base {
		base[i] = float64(i%100) + 1.0
	}
	exponent := 2.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = PowScalarF64(base, exponent)
	}
}

func BenchmarkPowF64_GoStdlib_1000(b *testing.B) {
	base := make([]float64, 1000)
	exponent := make([]float64, 1000)
	output := make([]float64, 1000)
	for i := range base {
		base[i] = float64(i%100) + 1.0
		exponent[i] = 2.0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range base {
			output[j] = math.Pow(base[j], exponent[j])
		}
	}
}
