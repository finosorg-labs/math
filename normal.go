package math

/*
#include "normal.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// NormalPdfF64 computes the standard normal probability density function φ(x)
// for each element in a float64 slice.
// Returns a new slice containing the PDF values.
func NormalPdfF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_normal_pdf_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_pdf_f64 failed with code %d", ret)
	}

	return output, nil
}

// NormalPdfF64InPlace computes the standard normal PDF in-place for a float64 slice.
// The input slice is modified to contain the PDF values.
func NormalPdfF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_pdf_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_pdf_f64 failed with code %d", ret)
	}

	return nil
}

// NormalPdfF32 computes the standard normal PDF for each element in a float32 slice.
// Returns a new slice containing the PDF values.
func NormalPdfF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_normal_pdf_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_pdf_f32 failed with code %d", ret)
	}

	return output, nil
}

// NormalPdfF32InPlace computes the standard normal PDF in-place for a float32 slice.
// The input slice is modified to contain the PDF values.
func NormalPdfF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_pdf_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_pdf_f32 failed with code %d", ret)
	}

	return nil
}

// NormalCdfF64 computes the standard normal cumulative distribution function Φ(x)
// for each element in a float64 slice.
// Returns a new slice containing the CDF values.
func NormalCdfF64(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}

	output := make([]float64, len(input))

	ret := C.fc_math_normal_cdf_f64(
		(*C.double)(unsafe.Pointer(&input[0])),
		(*C.double)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_cdf_f64 failed with code %d", ret)
	}

	return output, nil
}

// NormalCdfF64InPlace computes the standard normal CDF in-place for a float64 slice.
// The input slice is modified to contain the CDF values.
func NormalCdfF64InPlace(data []float64) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_cdf_f64(
		(*C.double)(unsafe.Pointer(&data[0])),
		(*C.double)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_cdf_f64 failed with code %d", ret)
	}

	return nil
}

// NormalCdfF32 computes the standard normal CDF for each element in a float32 slice.
// Returns a new slice containing the CDF values.
func NormalCdfF32(input []float32) ([]float32, error) {
	if len(input) == 0 {
		return []float32{}, nil
	}

	output := make([]float32, len(input))

	ret := C.fc_math_normal_cdf_f32(
		(*C.float)(unsafe.Pointer(&input[0])),
		(*C.float)(unsafe.Pointer(&output[0])),
		C.size_t(len(input)),
	)

	if ret != 0 {
		return nil, fmt.Errorf("fc_math_normal_cdf_f32 failed with code %d", ret)
	}

	return output, nil
}

// NormalCdfF32InPlace computes the standard normal CDF in-place for a float32 slice.
// The input slice is modified to contain the CDF values.
func NormalCdfF32InPlace(data []float32) error {
	if len(data) == 0 {
		return nil
	}

	ret := C.fc_math_normal_cdf_f32(
		(*C.float)(unsafe.Pointer(&data[0])),
		(*C.float)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
	)

	if ret != 0 {
		return fmt.Errorf("fc_math_normal_cdf_f32 failed with code %d", ret)
	}

	return nil
}
