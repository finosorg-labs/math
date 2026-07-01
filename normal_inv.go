package math

/*
#include "normal_inv.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// NormalInvF64 computes the inverse standard normal cumulative distribution function (quantile function)
// Φ⁻¹(p) for each probability p in a float64 slice using Acklam's algorithm (high accuracy).
// Returns a new slice containing the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_normal_inv_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_inv_f64 failed with code %d", ret)
	}

	return output, nil
}

// NormalInvBsmF64 computes the inverse standard normal CDF using original Beasley-Springer-Moro algorithm.
// For higher accuracy, use NormalInvF64 which implements Acklam's improved version.
// Returns a new slice containing the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvBsmF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_normal_inv_bsm_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_inv_bsm_f64 failed with code %d", ret)
	}

	return output, nil
}

// NormalInvF64InPlace computes the inverse standard normal CDF in-place for a float64 slice.
// The input slice is modified to contain the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_inv_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_inv_f64 failed with code %d", ret)
	}

	return nil
}

// NormalInvBsmF64InPlace computes the inverse standard normal CDF in-place using BSM algorithm.
// The input slice is modified to contain the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvBsmF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_inv_bsm_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_inv_bsm_f64 failed with code %d", ret)
	}

	return nil
}

// NormalInvF32 computes the inverse standard normal CDF for each element in a float32 slice.
// Returns a new slice containing the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_normal_inv_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_inv_f32 failed with code %d", ret)
	}

	return output, nil
}

// NormalInvBsmF32 computes the inverse standard normal CDF using BSM algorithm for float32.
// Returns a new slice containing the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvBsmF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_normal_inv_bsm_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_inv_bsm_f32 failed with code %d", ret)
	}

	return output, nil
}

// NormalInvF32InPlace computes the inverse standard normal CDF in-place for a float32 slice.
// The input slice is modified to contain the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_inv_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_inv_f32 failed with code %d", ret)
	}

	return nil
}

// NormalInvBsmF32InPlace computes the inverse standard normal CDF in-place using BSM algorithm.
// The input slice is modified to contain the quantile values.
// Input probabilities must be in the range (0, 1) exclusive.
func NormalInvBsmF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_inv_bsm_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_inv_bsm_f32 failed with code %d", ret)
	}

	return nil
}
