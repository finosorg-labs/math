#include "cumsum.h"
#include "bench_framework.h"
#include <stdlib.h>
#include <string.h>

typedef struct {
    double* input;
    double* output;
    size_t n;
} bench_cumsum_ctx_t;

static void bench_cumsum_f64_func(void* user_data) {
    bench_cumsum_ctx_t* ctx = (bench_cumsum_ctx_t*)user_data;
    fc_math_cumsum_f64(ctx->input, ctx->output, ctx->n);
}

static void bench_cumsum_f32_func(void* user_data) {
    bench_cumsum_ctx_t* ctx = (bench_cumsum_ctx_t*)user_data;
    fc_math_cumsum_f32((float*)ctx->input, (float*)ctx->output, ctx->n);
}

static void run_bench_cumsum_f64(const char* name, size_t n) {
    bench_cumsum_ctx_t ctx;
    ctx.n = n;
    ctx.input = (double*)malloc(n * sizeof(double));
    ctx.output = (double*)malloc(n * sizeof(double));

    for (size_t i = 0; i < n; i++) {
        ctx.input[i] = (double)(i % 100) / 100.0;
    }

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(double);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_cumsum_f64_func, &ctx, &result);

    free(ctx.input);
    free(ctx.output);
}

static void run_bench_cumsum_f32(const char* name, size_t n) {
    bench_cumsum_ctx_t ctx;
    ctx.n = n;
    ctx.input = (double*)malloc(n * sizeof(float));
    ctx.output = (double*)malloc(n * sizeof(float));

    for (size_t i = 0; i < n; i++) {
        ((float*)ctx.input)[i] = (float)(i % 100) / 100.0f;
    }

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(float);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_cumsum_f32_func, &ctx, &result);

    free(ctx.input);
    free(ctx.output);
}

void bench_cumsum_run(void) {
    run_bench_cumsum_f64("cumsum_f64[n=100]", 100);
    run_bench_cumsum_f64("cumsum_f64[n=1000]", 1000);
    run_bench_cumsum_f64("cumsum_f64[n=10000]", 10000);
    run_bench_cumsum_f64("cumsum_f64[n=100000]", 100000);
    run_bench_cumsum_f32("cumsum_f32[n=1000]", 1000);
}
