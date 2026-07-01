/**
 * @file test_math.c
 * @brief math module test entry point
 *
 * This file serves as the main test registration point for the math module.
 * Individual test modules are in separate files:
 */

#include "test_framework.h"

/* External test registration functions from sub-modules */
extern void register_cumsum_tests(void);
extern void register_cummax_tests(void);
extern void register_exp_tests(void);
extern void register_log_tests(void);
extern void register_erf_tests(void);
extern void register_normal_tests(void);
extern void register_normal_inv_tests(void);

/* Entry point for math tests */
void register_math_tests(void) {
    /* Register all sub-module tests */
    register_cumsum_tests();
    register_cummax_tests();
    register_exp_tests();
    register_log_tests();
    register_erf_tests();
    register_normal_tests();
    register_normal_inv_tests();
}
