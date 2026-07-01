#include "../tests/test_framework.h"
#include "erf.h"
#include <math.h>
#include <float.h>

TEST(test_erf_f64_basic) {
    double input[] = {0.0, 1.0, 2.0, -1.0, 0.5};
    double output[5];
    double expected[5];

    for (int i = 0; i < 5; i++) {
        expected[i] = erf(input[i]);
    }

    int ret = fc_math_erf_f64(input, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-6, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_erf_f64_zero) {
    double input[] = {0.0};
    double output[1];

    int ret = fc_math_erf_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 0.0, 1e-15);
}

TEST(test_erf_f64_symmetry) {
    double input[] = {1.0, 2.0, 3.0};
    double input_neg[] = {-1.0, -2.0, -3.0};
    double output[3];
    double output_neg[3];

    fc_math_erf_f64(input, output, 3);
    fc_math_erf_f64(input_neg, output_neg, 3);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_DOUBLE_EQ(output[i], -output_neg[i], 1e-15);
    }
}

TEST(test_erf_f64_inf) {
    double input[] = {INFINITY, -INFINITY};
    double output[2];

    int ret = fc_math_erf_f64(input, output, 2);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 1.0, 1e-15);
    FC_TEST_ASSERT_DOUBLE_EQ(output[1], -1.0, 1e-15);
}

TEST(test_erf_f64_nan) {
    double input[] = {NAN};
    double output[1];

    int ret = fc_math_erf_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "erf(NaN) should be NaN");
}

TEST(test_erf_f64_null_input) {
    double output[1];
    int ret = fc_math_erf_f64(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_erf_f64_null_output) {
    double input[] = {1.0};
    int ret = fc_math_erf_f64(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_erf_f64_empty) {
    double input[] = {0.0};
    double output[1];
    int ret = fc_math_erf_f64(input, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_erf_f64_large) {
    const size_t n = 10000;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));
    double* expected = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = ((double)i / (double)n) * 6.0 - 3.0;
        expected[i] = erf(input[i]);
    }

    int ret = fc_math_erf_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-5, "Index %zu: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }

    free(input);
    free(output);
    free(expected);
}

void register_erf_tests(void) {
    RUN_TEST(test_erf_f64_basic);
    RUN_TEST(test_erf_f64_zero);
    RUN_TEST(test_erf_f64_symmetry);
    RUN_TEST(test_erf_f64_inf);
    RUN_TEST(test_erf_f64_nan);
    RUN_TEST(test_erf_f64_null_input);
    RUN_TEST(test_erf_f64_null_output);
    RUN_TEST(test_erf_f64_empty);
    RUN_TEST(test_erf_f64_large);
}
