package math

/*
#include "exp.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ExpF64 computes exp(x) for each element in a float64 slice.
// Returns a new slice containing the exponentials.
func ExpF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_exp_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_exp_f64 failed with code %d", ret)
	}

	return output, nil
}

// ExpF64InPlace computes exp(x) in-place for a float64 slice.
// The input slice is modified to contain the exponentials.
func ExpF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_exp_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_exp_f64 failed with code %d", ret)
	}

	return nil
}

// ExpF32 computes exp(x) for each element in a float32 slice.
// Returns a new slice containing the exponentials.
func ExpF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_exp_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_exp_f32 failed with code %d", ret)
	}

	return output, nil
}

// ExpF32InPlace computes exp(x) in-place for a float32 slice.
// The input slice is modified to contain the exponentials.
func ExpF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_exp_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_exp_f32 failed with code %d", ret)
	}

	return nil
}
