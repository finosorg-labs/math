package math

/*
#include "pow.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// PowF64 computes pow(base, exponent) for each pair of elements.
// Returns a new slice containing base[i]^exponent[i].
func PowF64(base, exponent []float64) ([]float64, error) {
	if len(base) != len(exponent) {
		return nil, fmt.Errorf("base and exponent slices must have the same length")
	}

	if len(base) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(base))

	ret := C.fc_math_pow_f64(
		(*C.double)(unsafe.Pointer(&base[0])),
		(*C.double)(unsafe.Pointer(&exponent[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_pow_f64 failed with code %d", ret)
	}

	return output, nil
}

// PowF64InPlace computes pow(base, exponent) in-place.
// The base slice is modified to contain base[i]^exponent[i].
func PowF64InPlace(base, exponent []float64) error {
	if len(base) != len(exponent) {
		return fmt.Errorf("base and exponent slices must have the same length")
	}

	if len(base) == 0 {
		return nil
	}

	ret := C.fc_math_pow_f64(
		(*C.double)(unsafe.Pointer(&base[0])),
		(*C.double)(unsafe.Pointer(&exponent[0])),
		(*C.double)(unsafe.Pointer(&base[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_pow_f64 failed with code %d", ret)
	}

	return nil
}

// PowScalarF64 computes pow(base[i], exponent) for each element with a scalar exponent.
// Returns a new slice containing base[i]^exponent.
func PowScalarF64(base []float64, exponent float64) ([]float64, error) {
	if len(base) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(base))

	ret := C.fc_math_pow_scalar_f64(
		(*C.double)(unsafe.Pointer(&base[0])),
		C.double(exponent),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_pow_scalar_f64 failed with code %d", ret)
	}

	return output, nil
}

// PowScalarF64InPlace computes pow(base[i], exponent) in-place with a scalar exponent.
// The base slice is modified to contain base[i]^exponent.
func PowScalarF64InPlace(base []float64, exponent float64) error {
	if len(base) == 0 {
		return nil
	}

	ret := C.fc_math_pow_scalar_f64(
		(*C.double)(unsafe.Pointer(&base[0])),
		C.double(exponent),
		(*C.double)(unsafe.Pointer(&base[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_pow_scalar_f64 failed with code %d", ret)
	}

	return nil
}

// PowF32 computes pow(base, exponent) for each pair of elements.
// Returns a new slice containing base[i]^exponent[i].
func PowF32(base, exponent []float32) ([]float32, error) {
	if len(base) != len(exponent) {
		return nil, fmt.Errorf("base and exponent slices must have the same length")
	}

	if len(base) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(base))

	ret := C.fc_math_pow_f32(
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&exponent[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_pow_f32 failed with code %d", ret)
	}

	return output, nil
}

// PowF32InPlace computes pow(base, exponent) in-place.
// The base slice is modified to contain base[i]^exponent[i].
func PowF32InPlace(base, exponent []float32) error {
	if len(base) != len(exponent) {
		return fmt.Errorf("base and exponent slices must have the same length")
	}

	if len(base) == 0 {
		return nil
	}

	ret := C.fc_math_pow_f32(
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&exponent[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_pow_f32 failed with code %d", ret)
	}

	return nil
}

// PowScalarF32 computes pow(base[i], exponent) for each element with a scalar exponent.
// Returns a new slice containing base[i]^exponent.
func PowScalarF32(base []float32, exponent float32) ([]float32, error) {
	if len(base) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(base))

	ret := C.fc_math_pow_scalar_f32(
		(*C.float)(unsafe.Pointer(&base[0])),
		C.float(exponent),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_pow_scalar_f32 failed with code %d", ret)
	}

	return output, nil
}

// PowScalarF32InPlace computes pow(base[i], exponent) in-place with a scalar exponent.
// The base slice is modified to contain base[i]^exponent.
func PowScalarF32InPlace(base []float32, exponent float32) error {
	if len(base) == 0 {
		return nil
	}

	ret := C.fc_math_pow_scalar_f32(
		(*C.float)(unsafe.Pointer(&base[0])),
		C.float(exponent),
		(*C.float)(unsafe.Pointer(&base[0])),
		C.size_t(len(base)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_pow_scalar_f32 failed with code %d", ret)
	}

	return nil
}
