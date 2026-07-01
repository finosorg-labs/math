#include "pow.h"
#include "bench_framework.h"
#include "mem_aligned.h"
#include <stdlib.h>
#include <string.h>

typedef struct {
    double* base;
    double* exponent;
    double* output;
    size_t n;
} bench_pow_ctx_t;

static void bench_pow_f64_func(void* user_data) {
    bench_pow_ctx_t* ctx = (bench_pow_ctx_t*)user_data;
    fc_math_pow_f64(ctx->base, ctx->exponent, ctx->output, ctx->n);
}

static void bench_pow_scalar_f64_func(void* user_data) {
    bench_pow_ctx_t* ctx = (bench_pow_ctx_t*)user_data;
    fc_math_pow_scalar_f64(ctx->base, ctx->exponent[0], ctx->output, ctx->n);
}

static void bench_pow_f32_func(void* user_data) {
    bench_pow_ctx_t* ctx = (bench_pow_ctx_t*)user_data;
    fc_math_pow_f32((float*)ctx->base, (float*)ctx->exponent, (float*)ctx->output, ctx->n);
}

static void run_bench_pow_f64(const char* name, size_t n) {
    bench_pow_ctx_t ctx;
    ctx.n = n;
    ctx.base = (double*)fc_aligned_alloc_double(n);
    ctx.exponent = (double*)fc_aligned_alloc_double(n);
    ctx.output = (double*)fc_aligned_alloc_double(n);

    for (size_t i = 0; i < n; i++) {
        ctx.base[i] = ((double)(i % 100) / 10.0) + 1.0;
        ctx.exponent[i] = 2.0;
    }

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(double);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_pow_f64_func, &ctx, &result);

    fc_aligned_free(ctx.base);
    fc_aligned_free(ctx.exponent);
    fc_aligned_free(ctx.output);
}

static void run_bench_pow_scalar_f64(const char* name, size_t n) {
    bench_pow_ctx_t ctx;
    ctx.n = n;
    ctx.base = (double*)fc_aligned_alloc_double(n);
    ctx.exponent = (double*)fc_aligned_alloc_double(1);
    ctx.output = (double*)fc_aligned_alloc_double(n);

    for (size_t i = 0; i < n; i++) {
        ctx.base[i] = ((double)(i % 100) / 10.0) + 1.0;
    }
    ctx.exponent[0] = 2.0;

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(double);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_pow_scalar_f64_func, &ctx, &result);

    fc_aligned_free(ctx.base);
    fc_aligned_free(ctx.exponent);
    fc_aligned_free(ctx.output);
}

static void run_bench_pow_f32(const char* name, size_t n) {
    bench_pow_ctx_t ctx;
    ctx.n = n;
    ctx.base = (double*)fc_aligned_alloc_float(n);
    ctx.exponent = (double*)fc_aligned_alloc_float(n);
    ctx.output = (double*)fc_aligned_alloc_float(n);

    for (size_t i = 0; i < n; i++) {
        ((float*)ctx.base)[i] = ((float)(i % 100) / 10.0f) + 1.0f;
        ((float*)ctx.exponent)[i] = 2.0f;
    }

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(float);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_pow_f32_func, &ctx, &result);

    fc_aligned_free(ctx.base);
    fc_aligned_free(ctx.exponent);
    fc_aligned_free(ctx.output);
}

void bench_pow_run(void) {
    run_bench_pow_f64("pow_f64[n=100]", 100);
    run_bench_pow_f64("pow_f64[n=1000]", 1000);
    run_bench_pow_f64("pow_f64[n=10000]", 10000);
    run_bench_pow_scalar_f64("pow_scalar_f64[n=10000]", 10000);
    run_bench_pow_f32("pow_f32[n=1000]", 1000);
}
