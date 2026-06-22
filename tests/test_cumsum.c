#include "cumsum.h"
#include "test_framework.h"
#include <math.h>

TEST(cumsum_f64_basic) {
    double input[] = {1.0, 2.0, 3.0, 4.0, 5.0};
    double output[5];
    double expected[] = {1.0, 3.0, 6.0, 10.0, 15.0};

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        ASSERT_TRUE(fabs(output[i] - expected[i]) < 1e-10);
    }
}

TEST(cumsum_f32_basic) {
    float input[] = {1.0f, 2.0f, 3.0f, 4.0f, 5.0f};
    float output[5];
    float expected[] = {1.0f, 3.0f, 6.0f, 10.0f, 15.0f};

    int ret = fc_math_cumsum_f32(input, output, 5);
    ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        ASSERT_TRUE(fabs(output[i] - expected[i]) < 1e-6);
    }
}

TEST(cumsum_f64_negative) {
    double input[] = {-1.0, -2.0, -3.0, -4.0, -5.0};
    double output[5];
    double expected[] = {-1.0, -3.0, -6.0, -10.0, -15.0};

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        ASSERT_TRUE(fabs(output[i] - expected[i]) < 1e-10);
    }
}

TEST(cumsum_f64_mixed) {
    double input[] = {1.0, -2.0, 3.0, -4.0, 5.0};
    double output[5];
    double expected[] = {1.0, -1.0, 2.0, -2.0, 3.0};

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        ASSERT_TRUE(fabs(output[i] - expected[i]) < 1e-10);
    }
}

TEST(cumsum_f64_nan) {
    double input[] = {1.0, 2.0, NAN, 4.0, 5.0};
    double output[5];

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    ASSERT_TRUE(fabs(output[0] - 1.0) < 1e-10);
    ASSERT_TRUE(fabs(output[1] - 3.0) < 1e-10);
    ASSERT_TRUE(isnan(output[2]));
    ASSERT_TRUE(isnan(output[3]));
    ASSERT_TRUE(isnan(output[4]));
}

TEST(cumsum_f64_inf) {
    double input[] = {1.0, INFINITY, 3.0, 4.0, 5.0};
    double output[5];

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    ASSERT_TRUE(fabs(output[0] - 1.0) < 1e-10);
    ASSERT_TRUE(isinf(output[1]) && output[1] > 0);
    ASSERT_TRUE(isinf(output[2]) && output[2] > 0);
    ASSERT_TRUE(isinf(output[3]) && output[3] > 0);
    ASSERT_TRUE(isinf(output[4]) && output[4] > 0);
}

TEST(cumsum_f64_zero) {
    double input[] = {0.0, 0.0, 0.0, 0.0, 0.0};
    double output[5];
    double expected[] = {0.0, 0.0, 0.0, 0.0, 0.0};

    int ret = fc_math_cumsum_f64(input, output, 5);
    ASSERT_EQ(ret, 0);

    for (int i = 0; i < 5; i++) {
        ASSERT_TRUE(fabs(output[i] - expected[i]) < 1e-10);
    }
}

TEST(cumsum_f64_single) {
    double input[] = {42.0};
    double output[1];
    double expected[] = {42.0};

    int ret = fc_math_cumsum_f64(input, output, 1);
    ASSERT_EQ(ret, 0);

    ASSERT_TRUE(fabs(output[0] - expected[0]) < 1e-10);
}

TEST(cumsum_f64_large) {
    const size_t n = 1000;
    double input[1000];
    double output[1000];

    for (size_t i = 0; i < n; i++) {
        input[i] = 1.0;
    }

    int ret = fc_math_cumsum_f64(input, output, n);
    ASSERT_EQ(ret, 0);

    for (size_t i = 0; i < n; i++) {
        ASSERT_TRUE(fabs(output[i] - (double)(i + 1)) < 1e-9);
    }
}

TEST(cumsum_f64_null) {
    double input[] = {1.0, 2.0, 3.0};
    double output[3];

    ASSERT_EQ(fc_math_cumsum_f64(NULL, output, 3), -1);
    ASSERT_EQ(fc_math_cumsum_f64(input, NULL, 3), -1);
}

TEST(cumsum_f64_empty) {
    double input[] = {1.0};
    double output[1];

    int ret = fc_math_cumsum_f64(input, output, 0);
    ASSERT_EQ(ret, 0);
}

void register_cumsum_tests(void) {
    RUN_TEST(cumsum_f64_basic);
    RUN_TEST(cumsum_f32_basic);
    RUN_TEST(cumsum_f64_negative);
    RUN_TEST(cumsum_f64_mixed);
    RUN_TEST(cumsum_f64_nan);
    RUN_TEST(cumsum_f64_inf);
    RUN_TEST(cumsum_f64_zero);
    RUN_TEST(cumsum_f64_single);
    RUN_TEST(cumsum_f64_large);
    RUN_TEST(cumsum_f64_null);
    RUN_TEST(cumsum_f64_empty);
}
