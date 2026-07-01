#include "../tests/test_framework.h"
#include "pow.h"
#include <math.h>
#include <float.h>

TEST(test_pow_f64_basic) {
    double base[] = {2.0, 3.0, 10.0, 0.5, 2.0};
    double exponent[] = {3.0, 2.0, 1.0, -1.0, 0.5};
    double output[5];
    double expected[5];

    for (int i = 0; i < 5; i++) {
        expected[i] = pow(base[i], exponent[i]);
    }

    int ret = fc_math_pow_f64(base, exponent, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_pow_f64_zero_exponent) {
    double base[] = {2.0, -5.0, 0.0, NAN};
    double exponent[] = {0.0, 0.0, 0.0, 0.0};
    double output[4];

    int ret = fc_math_pow_f64(base, exponent, output, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        FC_TEST_ASSERT_MSG(output[i] == 1.0, "Index %d: pow(x, 0) should be 1.0", i);
    }
}

TEST(test_pow_f64_base_one) {
    double base[] = {1.0, 1.0, 1.0};
    double exponent[] = {5.0, -3.0, NAN};
    double output[3];

    int ret = fc_math_pow_f64(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(output[i] == 1.0, "Index %d: pow(1, y) should be 1.0", i);
    }
}

TEST(test_pow_f64_zero_base) {
    double base[] = {0.0, 0.0, 0.0};
    double exponent[] = {2.0, -2.0, 0.0};
    double output[3];

    int ret = fc_math_pow_f64(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(output[0] == 0.0, "pow(0, 2) should be 0");
    FC_TEST_ASSERT_MSG(isinf(output[1]) && output[1] > 0, "pow(0, -2) should be +Inf");
    FC_TEST_ASSERT_MSG(output[2] == 1.0, "pow(0, 0) should be 1");
}

TEST(test_pow_f64_negative_base) {
    double base[] = {-2.0};
    double exponent[] = {0.5};
    double output[1];

    int ret = fc_math_pow_f64(base, exponent, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "pow(-2, 0.5) should be NaN");
}

TEST(test_pow_f64_inf) {
    double base[] = {INFINITY, INFINITY, -INFINITY};
    double exponent[] = {2.0, -2.0, 2.0};
    double output[3];

    int ret = fc_math_pow_f64(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(isinf(output[0]) && output[0] > 0, "pow(+Inf, 2) should be +Inf");
    FC_TEST_ASSERT_MSG(output[1] == 0.0, "pow(+Inf, -2) should be 0");
}

TEST(test_pow_f64_nan) {
    double base[] = {NAN, 2.0};
    double exponent[] = {2.0, NAN};
    double output[2];

    int ret = fc_math_pow_f64(base, exponent, output, 2);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(isnan(output[0]), "pow(NaN, 2) should be NaN");
    FC_TEST_ASSERT_MSG(isnan(output[1]), "pow(2, NaN) should be NaN");
}

TEST(test_pow_f64_null_input) {
    double exponent[] = {2.0};
    double output[1];
    int ret = fc_math_pow_f64(NULL, exponent, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_pow_f64_null_exponent) {
    double base[] = {2.0};
    double output[1];
    int ret = fc_math_pow_f64(base, NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_pow_f64_null_output) {
    double base[] = {2.0};
    double exponent[] = {2.0};
    int ret = fc_math_pow_f64(base, exponent, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_pow_f64_empty) {
    double base[] = {2.0};
    double exponent[] = {2.0};
    double output[1];
    int ret = fc_math_pow_f64(base, exponent, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_pow_scalar_f64_basic) {
    double base[] = {2.0, 3.0, 4.0, 5.0};
    double exponent = 2.0;
    double output[4];
    double expected[4];

    for (int i = 0; i < 4; i++) {
        expected[i] = pow(base[i], exponent);
    }

    int ret = fc_math_pow_scalar_f64(base, exponent, output, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-4, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_pow_scalar_f64_zero_exponent) {
    double base[] = {2.0, -5.0, 0.0};
    double exponent = 0.0;
    double output[3];

    int ret = fc_math_pow_scalar_f64(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(output[i] == 1.0, "Index %d: pow(x, 0) should be 1.0", i);
    }
}

TEST(test_pow_f32_basic) {
    float base[] = {2.0f, 3.0f, 10.0f};
    float exponent[] = {3.0f, 2.0f, 1.0f};
    float output[3];
    float expected[3];

    for (int i = 0; i < 3; i++) {
        expected[i] = powf(base[i], exponent[i]);
    }

    int ret = fc_math_pow_f32(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        float abs_err = fabsf(output[i] - expected[i]);
        float rel_err = fabsf(expected[i]) > 1e-7f ? abs_err / fabsf(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 1e-3f, "Index %d: abs_err = %e, rel_err = %e", i, abs_err, rel_err);
    }
}

TEST(test_pow_f32_zero_exponent) {
    float base[] = {2.0f, -5.0f, 0.0f};
    float exponent[] = {0.0f, 0.0f, 0.0f};
    float output[3];

    int ret = fc_math_pow_f32(base, exponent, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        FC_TEST_ASSERT_MSG(output[i] == 1.0f, "Index %d: pow(x, 0) should be 1.0", i);
    }
}

void register_pow_tests(void) {
    RUN_TEST(test_pow_f64_basic);
    RUN_TEST(test_pow_f64_zero_exponent);
    RUN_TEST(test_pow_f64_base_one);
    RUN_TEST(test_pow_f64_zero_base);
    RUN_TEST(test_pow_f64_negative_base);
    RUN_TEST(test_pow_f64_inf);
    RUN_TEST(test_pow_f64_nan);
    RUN_TEST(test_pow_f64_null_input);
    RUN_TEST(test_pow_f64_null_exponent);
    RUN_TEST(test_pow_f64_null_output);
    RUN_TEST(test_pow_f64_empty);
    RUN_TEST(test_pow_scalar_f64_basic);
    RUN_TEST(test_pow_scalar_f64_zero_exponent);
    RUN_TEST(test_pow_f32_basic);
    RUN_TEST(test_pow_f32_zero_exponent);
}
