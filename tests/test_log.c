#include "../tests/test_framework.h"
#include "log.h"
#include <math.h>
#include <float.h>

TEST(test_log_f64_basic) {
    double input[] = {1.0, 2.0, 10.0, M_E, 0.5};
    double output[5];
    double expected[5];

    for (int i = 0; i < 5; i++) {
        expected[i] = log(input[i]);
    }

    int ret = fc_math_log_f64(input, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 5e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_log_f64_one) {
    double input[] = {1.0};
    double output[1];

    int ret = fc_math_log_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(fabs(output[0]) < 1e-15, "log(1) = %e, expected 0", output[0]);
}

TEST(test_log_f64_e) {
    double input[] = {M_E};
    double output[1];

    int ret = fc_math_log_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 1.0, 1e-15);
}

TEST(test_log_f64_zero) {
    double input[] = {0.0};
    double output[1];

    int ret = fc_math_log_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isinf(output[0]) && output[0] < 0, "log(0) should be -Inf");
}

TEST(test_log_f64_negative) {
    double input[] = {-1.0, -2.0, -10.0};
    double output[3];

    int ret = fc_math_log_f64(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(isnan(output[i]), "log(negative) should be NaN at index %d", i);
    }
}

TEST(test_log_f64_inf) {
    double input[] = {INFINITY};
    double output[1];

    int ret = fc_math_log_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isinf(output[0]) && output[0] > 0, "log(+Inf) should be +Inf");
}

TEST(test_log_f64_nan) {
    double input[] = {NAN};
    double output[1];

    int ret = fc_math_log_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "log(NaN) should be NaN");
}

TEST(test_log_f64_null_input) {
    double output[1];
    int ret = fc_math_log_f64(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_log_f64_null_output) {
    double input[] = {1.0};
    int ret = fc_math_log_f64(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_log_f64_empty) {
    double input[] = {0.0};
    double output[1];
    int ret = fc_math_log_f64(input, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_log_f64_large) {
    const size_t n = 10000;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = (double)(i + 1) / 10.0;
    }

    int ret = fc_math_log_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        double expected = log(input[i]);
        double abs_err = fabs(output[i] - expected);
        double rel_err = fabs(expected) > 1e-15 ? abs_err / fabs(expected) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 5e-4, "Index %zu: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }

    free(input);
    free(output);
}

TEST(test_log_f64_small_values) {
    double input[] = {1e-10, 1e-100, 1e-300};
    double output[3];

    int ret = fc_math_log_f64(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        double expected = log(input[i]);
        double abs_err = fabs(output[i] - expected);
        FC_TEST_ASSERT_MSG(abs_err < 1e-13, "Index %d: abs_err = %e", i, abs_err);
    }
}

TEST(test_log_f32_basic) {
    float input[] = {1.0f, 2.0f, 10.0f, 0.5f};
    float output[4];
    float expected[4];

    for (int i = 0; i < 4; i++) {
        expected[i] = logf(input[i]);
    }

    int ret = fc_math_log_f32(input, output, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        float abs_err = fabsf(output[i] - expected[i]);
        float rel_err = fabsf(expected[i]) > 1e-7f ? abs_err / fabsf(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-6f, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_log_f64_inplace) {
    double data[] = {1.0, 2.0, 10.0, M_E};
    double expected[4];

    for (int i = 0; i < 4; i++) {
        expected[i] = log(data[i]);
    }

    int ret = fc_math_log_f64(data, data, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        double abs_err = fabs(data[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

void register_log_tests(void) {
    RUN_TEST(test_log_f64_basic);
    RUN_TEST(test_log_f64_one);
    RUN_TEST(test_log_f64_e);
    RUN_TEST(test_log_f64_zero);
    RUN_TEST(test_log_f64_negative);
    RUN_TEST(test_log_f64_inf);
    RUN_TEST(test_log_f64_nan);
    RUN_TEST(test_log_f64_null_input);
    RUN_TEST(test_log_f64_null_output);
    RUN_TEST(test_log_f64_empty);
    RUN_TEST(test_log_f64_large);
    RUN_TEST(test_log_f64_small_values);
    RUN_TEST(test_log_f32_basic);
    RUN_TEST(test_log_f64_inplace);
}
