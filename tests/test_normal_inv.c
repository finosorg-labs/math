#include "../tests/test_framework.h"
#include "normal_inv.h"
#include <math.h>
#include <float.h>

TEST(test_normal_inv_f64_basic) {
    double input[] = {0.5, 0.8413447460685429, 0.1586552539314571, 0.9772498680518208, 0.0227501319481792};
    double expected[] = {0.0, 1.0, -1.0, 2.0, -2.0};
    double output[5];

    int ret = fc_math_normal_inv_f64(input, output, 5);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        FC_TEST_ASSERT_MSG(abs_err < 1e-6, "Index %d: expected = %f, got = %f, abs_err = %e",
                          i, expected[i], output[i], abs_err);
    }
}

TEST(test_normal_inv_f64_half) {
    double input[] = {0.5};
    double output[1];

    int ret = fc_math_normal_inv_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_DOUBLE_EQ(output[0], 0.0, 1e-15);
}

TEST(test_normal_inv_f64_symmetry) {
    double input[] = {0.1, 0.2, 0.3, 0.4};
    double input_complement[] = {0.9, 0.8, 0.7, 0.6};
    double output[4];
    double output_complement[4];

    fc_math_normal_inv_f64(input, output, 4);
    fc_math_normal_inv_f64(input_complement, output_complement, 4);

    for (int i = 0; i < 4; i++) {
        double sum = output[i] + output_complement[i];
        FC_TEST_ASSERT_MSG(fabs(sum) < 1e-9, "Index %d: symmetry failed, sum = %e", i, sum);
    }
}

TEST(test_normal_inv_f64_extreme_tails) {
    double input[] = {0.001, 0.999, 0.0001, 0.9999};
    double output[4];

    int ret = fc_math_normal_inv_f64(input, output, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    FC_TEST_ASSERT_MSG(output[0] < -3.0 && output[0] > -3.1, "0.001 quantile should be around -3.09");
    FC_TEST_ASSERT_MSG(output[1] > 3.0 && output[1] < 3.1, "0.999 quantile should be around 3.09");
    FC_TEST_ASSERT_MSG(output[2] < -3.7 && output[2] > -3.8, "0.0001 quantile should be around -3.72");
    FC_TEST_ASSERT_MSG(output[3] > 3.7 && output[3] < 3.8, "0.9999 quantile should be around 3.72");
}

TEST(test_normal_inv_f64_out_of_range) {
    double input[] = {0.0, 1.0, -0.5, 1.5};
    double output[4];

    int ret = fc_math_normal_inv_f64(input, output, 4);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 4; i++) {
        FC_TEST_ASSERT_MSG(isnan(output[i]), "Index %d: out-of-range input should produce NaN", i);
    }
}

TEST(test_normal_inv_f64_nan) {
    double input[] = {NAN};
    double output[1];

    int ret = fc_math_normal_inv_f64(input, output, 1);
    FC_TEST_ASSERT_EQ(ret, 0);
    FC_TEST_ASSERT_MSG(isnan(output[0]), "inv(NaN) should be NaN");
}

TEST(test_normal_inv_f64_null_input) {
    double output[1];
    int ret = fc_math_normal_inv_f64(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_normal_inv_f64_null_output) {
    double input[] = {0.5};
    int ret = fc_math_normal_inv_f64(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_normal_inv_f64_empty) {
    double input[] = {0.5};
    double output[1];
    int ret = fc_math_normal_inv_f64(input, output, 0);
    FC_TEST_ASSERT_EQ(ret, 0);
}

TEST(test_normal_inv_f64_known_values) {
    double input[] = {0.025, 0.05, 0.1, 0.25, 0.5, 0.75, 0.9, 0.95, 0.975};
    double expected[] = {-1.95996398454005, -1.64485362695147, -1.28155156554460,
                        -0.67448975019608, 0.0, 0.67448975019608,
                        1.28155156554460, 1.64485362695147, 1.95996398454005};
    double output[9];

    int ret = fc_math_normal_inv_f64(input, output, 9);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 9; i++) {
        double abs_err = fabs(output[i] - expected[i]);
        double rel_err = fabs(expected[i]) > 1e-15 ? abs_err / fabs(expected[i]) : abs_err;
        FC_TEST_ASSERT_MSG(rel_err < 2e-9, "Index %d: expected = %.15f, got = %.15f, rel_err = %e",
                          i, expected[i], output[i], rel_err);
    }
}

TEST(test_normal_inv_f64_central_region) {
    const size_t n = 100;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = 0.08 + ((double)i / (double)(n - 1)) * (0.92 - 0.08);
    }

    int ret = fc_math_normal_inv_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        FC_TEST_ASSERT_MSG(!isnan(output[i]) && isfinite(output[i]),
                          "Index %zu: central region should produce finite values", i);
        FC_TEST_ASSERT_MSG(output[i] >= -1.5 && output[i] <= 1.5,
                          "Index %zu: central region quantiles should be in [-1.5, 1.5]", i);
    }

    free(input);
    free(output);
}

TEST(test_normal_inv_f64_large) {
    const size_t n = 10000;
    double* input = malloc(n * sizeof(double));
    double* output = malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        input[i] = 0.001 + ((double)i / (double)(n - 1)) * (0.999 - 0.001);
    }

    int ret = fc_math_normal_inv_f64(input, output, n);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (size_t i = 1; i < n; i++) {
        FC_TEST_ASSERT_MSG(output[i] > output[i-1],
                          "Index %zu: quantile function should be monotonically increasing", i);
    }

    free(input);
    free(output);
}

TEST(test_normal_inv_f32_basic) {
    float input[] = {0.5f, 0.8413447460685429f, 0.1586552539314571f};
    float expected[] = {0.0f, 1.0f, -1.0f};
    float output[3];

    int ret = fc_math_normal_inv_f32(input, output, 3);
    FC_TEST_ASSERT_EQ(ret, 0);

    for (int i = 0; i < 3; i++) {
        float abs_err = fabsf(output[i] - expected[i]);
        FC_TEST_ASSERT_MSG(abs_err < 1e-4f, "Index %d: expected = %f, got = %f, abs_err = %e",
                          i, expected[i], output[i], abs_err);
    }
}

TEST(test_normal_inv_f32_null_input) {
    float output[1];
    int ret = fc_math_normal_inv_f32(NULL, output, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

TEST(test_normal_inv_f32_null_output) {
    float input[] = {0.5f};
    int ret = fc_math_normal_inv_f32(input, NULL, 1);
    FC_TEST_ASSERT_EQ(ret, -1);
}

void register_normal_inv_tests(void) {
    RUN_TEST(test_normal_inv_f64_basic);
    RUN_TEST(test_normal_inv_f64_half);
    RUN_TEST(test_normal_inv_f64_symmetry);
    RUN_TEST(test_normal_inv_f64_extreme_tails);
    RUN_TEST(test_normal_inv_f64_out_of_range);
    RUN_TEST(test_normal_inv_f64_nan);
    RUN_TEST(test_normal_inv_f64_null_input);
    RUN_TEST(test_normal_inv_f64_null_output);
    RUN_TEST(test_normal_inv_f64_empty);
    RUN_TEST(test_normal_inv_f64_known_values);
    RUN_TEST(test_normal_inv_f64_central_region);
    RUN_TEST(test_normal_inv_f64_large);

    RUN_TEST(test_normal_inv_f32_basic);
    RUN_TEST(test_normal_inv_f32_null_input);
    RUN_TEST(test_normal_inv_f32_null_output);
}
