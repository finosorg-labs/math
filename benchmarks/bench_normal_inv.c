#include "normal_inv.h"
#include "bench_framework.h"
#include "mem_aligned.h"
#include <stdlib.h>
#include <string.h>

typedef struct {
    double* input;
    double* output;
    size_t n;
} bench_normal_inv_ctx_t;

static void bench_normal_inv_f64_func(void* user_data) {
    bench_normal_inv_ctx_t* ctx = (bench_normal_inv_ctx_t*)user_data;
    fc_math_normal_inv_f64(ctx->input, ctx->output, ctx->n);
}

static void run_bench_normal_inv_f64(const char* name, size_t n) {
    bench_normal_inv_ctx_t ctx;
    ctx.n = n;
    ctx.input = (double*)fc_aligned_alloc_double(n);
    ctx.output = (double*)fc_aligned_alloc_double(n);

    for (size_t i = 0; i < n; i++) {
        ctx.input[i] = 0.001 + ((double)(i % 1000) / 1000.0) * (0.999 - 0.001);
    }

    fc_bench_config_t config = FC_BENCH_CONFIG_DEFAULT;
    config.name = name;
    config.data_size = n * sizeof(double);
    config.min_time_ms = 200.0;
    config.warmup_ms = 50.0;

    fc_bench_result_t result;
    fc_bench_run(&config, bench_normal_inv_f64_func, &ctx, &result);

    fc_aligned_free(ctx.input);
    fc_aligned_free(ctx.output);
}

void bench_normal_inv_run(void) {
    run_bench_normal_inv_f64("normal_inv_f64[n=100]", 100);
    run_bench_normal_inv_f64("normal_inv_f64[n=1000]", 1000);
    run_bench_normal_inv_f64("normal_inv_f64[n=10000]", 10000);
    run_bench_normal_inv_f64("normal_inv_f64[n=100000]", 100000);
}
