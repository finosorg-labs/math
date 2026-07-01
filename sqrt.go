package math

/*
#include "sqrt.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// SqrtF64 computes sqrt(x) for each element in a float64 slice.
// Returns a new slice containing the square roots.
func SqrtF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_sqrt_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_sqrt_f64 failed with code %d", ret)
	}

	return output, nil
}

// SqrtF64InPlace computes sqrt(x) in-place for a float64 slice.
// The input slice is modified to contain the square roots.
func SqrtF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_sqrt_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_sqrt_f64 failed with code %d", ret)
	}

	return nil
}

// SqrtF32 computes sqrt(x) for each element in a float32 slice.
// Returns a new slice containing the square roots.
func SqrtF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_sqrt_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_sqrt_f32 failed with code %d", ret)
	}

	return output, nil
}

// SqrtF32InPlace computes sqrt(x) in-place for a float32 slice.
// The input slice is modified to contain the square roots.
func SqrtF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_sqrt_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_sqrt_f32 failed with code %d", ret)
	}

	return nil
}
