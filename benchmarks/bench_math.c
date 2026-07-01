/**
 * @file bench_math.c
 * @brief math module benchmark entry point
 *
 * This file serves as the main benchmark registration point for the math module.
 * Individual benchmark modules are in separate files:
 */

#include "bench_framework.h"
#include <simd_detect.h>
#include <stdio.h>

/* External benchmark functions from sub-modules */
extern void bench_cumsum_run(void);
extern void bench_cummax_run(void);
extern void bench_exp_run(void);
extern void bench_log_run(void);
extern void bench_erf_run(void);
extern void bench_normal_run(void);
extern void bench_normal_inv_run(void);

/* Entry point for math benchmarks */
void bench_math_run(void) {
    printf("\n");
    printf("============================================================\n");
    printf("  math Module Performance Benchmarks\n");
    printf("  SIMD level: %s\n", fc_simd_level_string(fc_get_simd_level()));
    printf("============================================================\n");

    /* Run all sub-module benchmarks */
    bench_cumsum_run();
    bench_cummax_run();
    bench_exp_run();
    bench_log_run();
    bench_erf_run();
    bench_normal_run();
    bench_normal_inv_run();

    printf("\n");
    printf("============================================================\n");
    printf("  math benchmarks complete\n");
    printf("============================================================\n");
}
