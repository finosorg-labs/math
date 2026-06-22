package math

/*
#include "cummax.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// CummaxF64 computes the cumulative maximum of a float64 slice.
// Returns a new slice containing the cumulative maximums.
func CummaxF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_cummax_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_cummax_f64 failed with code %d", ret)
	}

	return output, nil
}

// CummaxF64InPlace computes the cumulative maximum of a float64 slice in-place.
// The input slice is modified to contain the cumulative maximums.
func CummaxF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_cummax_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_cummax_f64 failed with code %d", ret)
	}

	return nil
}

// CummaxF32 computes the cumulative maximum of a float32 slice.
// Returns a new slice containing the cumulative maximums.
func CummaxF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_cummax_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_cummax_f32 failed with code %d", ret)
	}

	return output, nil
}

// CummaxF32InPlace computes the cumulative maximum of a float32 slice in-place.
// The input slice is modified to contain the cumulative maximums.
func CummaxF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_cummax_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_cummax_f32 failed with code %d", ret)
	}

	return nil
}
