#ifndef FC_MATH_NORMAL_H
#define FC_MATH_NORMAL_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/**
 * @brief Compute standard normal probability density function (PDF)
 *
 * Computes φ(x) = (1/√(2π)) * exp(-x²/2) for each element in the input array.
 * Uses SIMD-optimized exponential function for high performance.
 *
 * @param input Input array (standard normal variates)
 * @param output Output array (probability densities)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note Output array must be pre-allocated with size n
 * @note Special value handling:
 *       - φ(NaN) = NaN
 *       - φ(±Inf) = 0
 * @note Thread-safe (no shared state)
 *
 * Time complexity: O(n)
 * Space complexity: O(1) auxiliary space
 */
int fc_math_normal_pdf_f64(const double* input, double* output, size_t n);

/**
 * @brief Compute standard normal cumulative distribution function (CDF)
 *
 * Computes Φ(x) = (1/2) * [1 + erf(x/√2)] for each element in the input array.
 * Uses rational Chebyshev approximation for erf() achieving <1 ULP accuracy.
 *
 * @param input Input array (standard normal variates)
 * @param output Output array (cumulative probabilities)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note Output array must be pre-allocated with size n
 * @note Special value handling:
 *       - Φ(NaN) = NaN
 *       - Φ(+Inf) = 1.0
 *       - Φ(-Inf) = 0.0
 * @note Thread-safe (no shared state)
 *
 * Time complexity: O(n)
 * Space complexity: O(1) auxiliary space
 */
int fc_math_normal_cdf_f64(const double* input, double* output, size_t n);

/**
 * @brief Compute error function erf(x)
 *
 * Computes erf(x) = (2/√π) * ∫[0,x] exp(-t²) dt using rational Chebyshev
 * approximation. Achieves <1 ULP accuracy across the entire range.
 *
 * @param input Input array
 * @param output Output array (erf values)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note Output array must be pre-allocated with size n
 * @note Special value handling:
 *       - erf(NaN) = NaN
 *       - erf(+Inf) = 1.0
 *       - erf(-Inf) = -1.0
 *       - erf(0) = 0 exactly
 * @note Thread-safe (no shared state)
 *
 * Time complexity: O(n)
 * Space complexity: O(1) auxiliary space
 */
int fc_math_erf_f64(const double* input, double* output, size_t n);

/**
 * @brief Compute standard normal PDF for float32 array
 *
 * @param input Input array (standard normal variates)
 * @param output Output array (probability densities)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note See fc_math_normal_pdf_f64 for details
 */
int fc_math_normal_pdf_f32(const float* input, float* output, size_t n);

/**
 * @brief Compute standard normal CDF for float32 array
 *
 * @param input Input array (standard normal variates)
 * @param output Output array (cumulative probabilities)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note See fc_math_normal_cdf_f64 for details
 */
int fc_math_normal_cdf_f32(const float* input, float* output, size_t n);

/**
 * @brief Compute error function for float32 array
 *
 * @param input Input array
 * @param output Output array (erf values)
 * @param n Number of elements
 * @return 0 on success, -1 on error
 *
 * @note See fc_math_erf_f64 for details
 */
int fc_math_erf_f32(const float* input, float* output, size_t n);

#ifdef __cplusplus
}
#endif

#endif
