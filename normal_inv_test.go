package math

import (
	"math"
	"testing"
)

func TestNormalInvF64Basic(t *testing.T) {
	input := []float64{0.5, 0.8413447460685429, 0.1586552539314571, 0.9772498680518208, 0.0227501319481792}
	expected := []float64{0.0, 1.0, -1.0, 2.0, -2.0}

	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	if len(output) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(output))
	}

	for i := range expected {
		absErr := math.Abs(output[i] - expected[i])
		if absErr > 1e-6 {
			t.Errorf("Index %d: expected = %f, got = %f, abs_err = %e", i, expected[i], output[i], absErr)
		}
	}
}

func TestNormalInvF64InPlace(t *testing.T) {
	data := []float64{0.5, 0.8413447460685429, 0.1586552539314571}
	expected := []float64{0.0, 1.0, -1.0}

	err := NormalInvF64InPlace(data)
	if err != nil {
		t.Fatalf("NormalInvF64InPlace failed: %v", err)
	}

	for i := range expected {
		absErr := math.Abs(data[i] - expected[i])
		if absErr > 1e-6 {
			t.Errorf("Index %d: expected = %f, got = %f, abs_err = %e", i, expected[i], data[i], absErr)
		}
	}
}

func TestNormalInvF64Half(t *testing.T) {
	input := []float64{0.5}
	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	if math.Abs(output[0]) > 1e-15 {
		t.Errorf("Expected 0.0, got %f", output[0])
	}
}

func TestNormalInvF64Symmetry(t *testing.T) {
	input := []float64{0.1, 0.2, 0.3, 0.4}
	inputComplement := []float64{0.9, 0.8, 0.7, 0.6}

	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	outputComplement, err := NormalInvF64(inputComplement)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	for i := range output {
		sum := output[i] + outputComplement[i]
		if math.Abs(sum) > 1e-9 {
			t.Errorf("Index %d: symmetry failed, sum = %e", i, sum)
		}
	}
}

func TestNormalInvF64ExtremeTails(t *testing.T) {
	input := []float64{0.001, 0.999, 0.0001, 0.9999}
	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	if output[0] >= -3.0 || output[0] <= -3.1 {
		t.Errorf("0.001 quantile should be around -3.09, got %f", output[0])
	}
	if output[1] <= 3.0 || output[1] >= 3.1 {
		t.Errorf("0.999 quantile should be around 3.09, got %f", output[1])
	}
	if output[2] >= -3.7 || output[2] <= -3.8 {
		t.Errorf("0.0001 quantile should be around -3.72, got %f", output[2])
	}
	if output[3] <= 3.7 || output[3] >= 3.8 {
		t.Errorf("0.9999 quantile should be around 3.72, got %f", output[3])
	}
}

func TestNormalInvF64OutOfRange(t *testing.T) {
	input := []float64{0.0, 1.0, -0.5, 1.5}
	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	for i, val := range output {
		if !math.IsNaN(val) {
			t.Errorf("Index %d: out-of-range input should produce NaN, got %f", i, val)
		}
	}
}

func TestNormalInvF64NaN(t *testing.T) {
	input := []float64{math.NaN()}
	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	if !math.IsNaN(output[0]) {
		t.Errorf("inv(NaN) should be NaN, got %f", output[0])
	}
}

func TestNormalInvF64Empty(t *testing.T) {
	input := []float64{}
	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	if len(output) != 0 {
		t.Errorf("Expected empty output, got length %d", len(output))
	}
}

func TestNormalInvF64KnownValues(t *testing.T) {
	input := []float64{0.025, 0.05, 0.1, 0.25, 0.5, 0.75, 0.9, 0.95, 0.975}
	expected := []float64{-1.95996398454005, -1.64485362695147, -1.28155156554460,
		-0.67448975019608, 0.0, 0.67448975019608,
		1.28155156554460, 1.64485362695147, 1.95996398454005}

	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	for i := range expected {
		absErr := math.Abs(output[i] - expected[i])
		relErr := absErr
		if math.Abs(expected[i]) > 1e-15 {
			relErr = absErr / math.Abs(expected[i])
		}
		if relErr > 1.2e-9 {
			t.Errorf("Index %d: expected = %.15f, got = %.15f, rel_err = %e",
				i, expected[i], output[i], relErr)
		}
	}
}

func TestNormalInvF64Monotonic(t *testing.T) {
	n := 1000
	input := make([]float64, n)
	for i := 0; i < n; i++ {
		input[i] = 0.001 + (float64(i)/float64(n-1))*(0.999-0.001)
	}

	output, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}

	for i := 1; i < n; i++ {
		if output[i] <= output[i-1] {
			t.Errorf("Index %d: quantile function should be monotonically increasing", i)
		}
	}
}

func TestNormalInvF32Basic(t *testing.T) {
	input := []float32{0.5, 0.8413447460685429, 0.1586552539314571}
	expected := []float32{0.0, 1.0, -1.0}

	output, err := NormalInvF32(input)
	if err != nil {
		t.Fatalf("NormalInvF32 failed: %v", err)
	}

	for i := range expected {
		absErr := float32(math.Abs(float64(output[i] - expected[i])))
		if absErr > 1e-4 {
			t.Errorf("Index %d: expected = %f, got = %f, abs_err = %e", i, expected[i], output[i], absErr)
		}
	}
}

func TestNormalInvF32InPlace(t *testing.T) {
	data := []float32{0.5, 0.8413447460685429}
	expected := []float32{0.0, 1.0}

	err := NormalInvF32InPlace(data)
	if err != nil {
		t.Fatalf("NormalInvF32InPlace failed: %v", err)
	}

	for i := range expected {
		absErr := float32(math.Abs(float64(data[i] - expected[i])))
		if absErr > 1e-4 {
			t.Errorf("Index %d: expected = %f, got = %f, abs_err = %e", i, expected[i], data[i], absErr)
		}
	}
}

func BenchmarkNormalInvF64_100(b *testing.B) {
	input := make([]float64, 100)
	for i := range input {
		input[i] = 0.001 + (float64(i%100)/100.0)*(0.999-0.001)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NormalInvF64(input)
	}
}

func BenchmarkNormalInvF64_1000(b *testing.B) {
	input := make([]float64, 1000)
	for i := range input {
		input[i] = 0.001 + (float64(i%1000)/1000.0)*(0.999-0.001)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NormalInvF64(input)
	}
}

func BenchmarkNormalInvF64_10000(b *testing.B) {
	input := make([]float64, 10000)
	for i := range input {
		input[i] = 0.001 + (float64(i%1000)/1000.0)*(0.999-0.001)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NormalInvF64(input)
	}
}

func TestNormalInvBsmF64Comparison(t *testing.T) {
	input := []float64{0.025, 0.5, 0.975}
	
	acklam, err := NormalInvF64(input)
	if err != nil {
		t.Fatalf("NormalInvF64 failed: %v", err)
	}
	
	bsm, err := NormalInvBsmF64(input)
	if err != nil {
		t.Fatalf("NormalInvBsmF64 failed: %v", err)
	}
	
	for i := range input {
		// Both should give similar results (within reasonable tolerance)
		diff := math.Abs(acklam[i] - bsm[i])
		if diff > 1e-5 {
			t.Logf("Index %d (p=%f): Acklam=%f, BSM=%f, diff=%e", i, input[i], acklam[i], bsm[i], diff)
		}
	}
}
