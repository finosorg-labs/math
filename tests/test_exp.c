#include "../tests/test_framework.h"
#include "exp.h"
#include <math.h>
#include <float.h>

TEST(test_exp_f64_basic) {
    double input[] = {0.0, 1.0, 2.0, -1.0, 0.5};
    double output[5];
    double expected[5];

    for (int i = 0; i < 5; i++) {
        expected[i] = exp(input[i]);
    }

    int ret = fc_math_exp_f64(input, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_exp_f64_zero) {
    double input[] = {0.0};
    double output[1];

    int ret = fc_math_exp_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 1.0, 1e-15);
}

TEST(test_exp_f64_overflow) {
    double input[] = {710.0, 750.0, 1000.0};
    double output[3];

    int ret = fc_math_exp_f64(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(isinf(output[i]) && output[i] > 0, "Index %d should be +Inf", i);
    }
}

TEST(test_exp_f64_underflow) {
    double input[] = {-750.0, -800.0, -1000.0};
    double output[3];

    int ret = fc_math_exp_f64(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(output[i] == 0.0, "Index %d should be 0", i);
    }
}

TEST(test_exp_f64_inf) {
    double input[] = {INFINITY, -INFINITY};
    double output[2];

    int ret = fc_math_exp_f64(input, output, 2);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(isinf(output[0]) && output[0] > 0, "exp(+Inf) should be +Inf");
    FC_TEST_ASSERT_MSG(output[1] == 0.0, "exp(-Inf) should be 0");
}

TEST(test_exp_f64_nan) {
    double input[] = {NAN};
    double output[1];

    int ret = fc_math_exp_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "exp(NaN) should be NaN");
}

TEST(test_exp_f64_null_input) {
    double output[1];
    int ret = fc_math_exp_f64(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_exp_f64_null_output) {
    double input[] = {1.0};
    int ret = fc_math_exp_f64(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_exp_f64_empty) {
    double input[] = {0.0};
    double output[1];
    int ret = fc_math_exp_f64(input, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_exp_f64_large) {
    const size_t n = 10000;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = (double)(i % 100) / 10.0;
    }

    int ret = fc_math_exp_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        double expected = exp(input[i]);
        double abs_err = fabs(output[i] - expected);
        double rel_err = fabs(expected) > 1e-15 ? abs_err / fabs(expected) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %zu: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }

    free(input);
    free(output);
}

TEST(test_exp_f32_basic) {
    float input[] = {0.0f, 1.0f, 2.0f, -1.0f, 0.5f};
    float output[5];
    float expected[5];

    for (int i = 0; i < 5; i++) {
        expected[i] = expf(input[i]);
    }

    int ret = fc_math_exp_f32(input, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        float rel_err = fabsf((output[i] - expected[i]) / expected[i]);
        FC_TEST_ASSERT_MSG(rel_err < 1e-6f, "Index %d: rel_err = %e", i, rel_err);
    }
}

TEST(test_exp_f64_inplace) {
    double data[] = {0.0, 1.0, 2.0, -1.0};
    double expected[4];

    for (int i = 0; i < 4; i++) {
        expected[i] = exp(data[i]);
    }

    int ret = fc_math_exp_f64(data, data, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        double abs_err = fabs(data[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

void register_exp_tests(void) {
    RUN_TEST(test_exp_f64_basic);
    RUN_TEST(test_exp_f64_zero);
    RUN_TEST(test_exp_f64_overflow);
    RUN_TEST(test_exp_f64_underflow);
    RUN_TEST(test_exp_f64_inf);
    RUN_TEST(test_exp_f64_nan);
    RUN_TEST(test_exp_f64_null_input);
    RUN_TEST(test_exp_f64_null_output);
    RUN_TEST(test_exp_f64_empty);
    RUN_TEST(test_exp_f64_large);
    RUN_TEST(test_exp_f32_basic);
    RUN_TEST(test_exp_f64_inplace);
}
