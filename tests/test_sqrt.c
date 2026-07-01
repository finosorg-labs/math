#include "../tests/test_framework.h"
#include "sqrt.h"
#include <math.h>
#include <float.h>

TEST(test_sqrt_f64_basic) {
    double input[] = {0.0, 1.0, 4.0, 9.0, 16.0, 0.25, 2.0};
    double output[7];
    double expected[7];

    for (int i = 0; i < 7; i++) {
        expected[i] = sqrt(input[i]);
    }

    int ret = fc_math_sqrt_f64(input, output, 7);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 7; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        FC_TEST_ASSERT_MSG(abs_err < 1e-15, "Index %d: abs_err = %e", i, abs_err);
    }
}

TEST(test_sqrt_f64_zero) {
    double input[] = {0.0};
    double output[1];

    int ret = fc_math_sqrt_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 0.0, 1e-15);
}

TEST(test_sqrt_f64_negative) {
    double input[] = {-1.0, -4.0, -0.5};
    double output[3];

    int ret = fc_math_sqrt_f64(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(isnan(output[i]), "Index %d should be NaN", i);
    }
}

TEST(test_sqrt_f64_inf) {
    double input[] = {INFINITY, -INFINITY};
    double output[2];

    int ret = fc_math_sqrt_f64(input, output, 2);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(isinf(output[0]) && output[0] > 0, "sqrt(+Inf) should be +Inf");
    FC_TEST_ASSERT_MSG(isnan(output[1]), "sqrt(-Inf) should be NaN");
}

TEST(test_sqrt_f64_nan) {
    double input[] = {NAN};
    double output[1];

    int ret = fc_math_sqrt_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "sqrt(NaN) should be NaN");
}

TEST(test_sqrt_f64_null_input) {
    double output[1];
    int ret = fc_math_sqrt_f64(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_sqrt_f64_null_output) {
    double input[] = {1.0};
    int ret = fc_math_sqrt_f64(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_sqrt_f64_empty) {
    double input[] = {0.0};
    double output[1];
    int ret = fc_math_sqrt_f64(input, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_sqrt_f64_large) {
    const size_t n = 10000;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));
    double* expected = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = (double)i * 0.1;
        expected[i] = sqrt(input[i]);
    }

    int ret = fc_math_sqrt_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        FC_TEST_ASSERT_MSG(abs_err < 1e-14, "Index %zu: abs_err = %e", i, abs_err);
    }

    free(input);
    free(output);
    free(expected);
}

TEST(test_sqrt_f32_basic) {
    float input[] = {0.0f, 1.0f, 4.0f, 9.0f, 16.0f, 0.25f};
    float output[6];
    float expected[6];

    for (int i = 0; i < 6; i++) {
        expected[i] = sqrtf(input[i]);
    }

    int ret = fc_math_sqrt_f32(input, output, 6);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 6; i++) {
        float abs_err = fabsf(output[i] - expected[i]);
        FC_TEST_ASSERT_MSG(abs_err < 1e-6f, "Index %d: abs_err = %e", i, abs_err);
    }
}

TEST(test_sqrt_f32_negative) {
    float input[] = {-1.0f, -4.0f};
    float output[2];

    int ret = fc_math_sqrt_f32(input, output, 2);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 2; i++) {
        FC_TEST_ASSERT_MSG(isnan(output[i]), "Index %d should be NaN", i);
    }
}

void register_sqrt_tests(void) {
    RUN_TEST(test_sqrt_f64_basic);
    RUN_TEST(test_sqrt_f64_zero);
    RUN_TEST(test_sqrt_f64_negative);
    RUN_TEST(test_sqrt_f64_inf);
    RUN_TEST(test_sqrt_f64_nan);
    RUN_TEST(test_sqrt_f64_null_input);
    RUN_TEST(test_sqrt_f64_null_output);
    RUN_TEST(test_sqrt_f64_empty);
    RUN_TEST(test_sqrt_f64_large);
    RUN_TEST(test_sqrt_f32_basic);
    RUN_TEST(test_sqrt_f32_negative);
}
