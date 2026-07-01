#
# benchmarks/CMakeLists.txt - Benchmark configuration for math module
#

# Benchmark framework library
add_library(bench_framework OBJECT
    bench_framework.c
)

target_include_directories(bench_framework
    PUBLIC
        ${CMAKE_CURRENT_SOURCE_DIR}
        ${CMAKE_SOURCE_DIR}/include
        ${CMAKE_SOURCE_DIR}/modules/platform/include
)

target_link_libraries(bench_framework
    PUBLIC
        finkit_math_static_base
        finkit::platform
      finkit::ds
)

# Benchmark executable
add_executable(math_benchmarks
    bench_main.c
    bench_math.c
    bench_cumsum.c
    bench_cummax.c
    bench_exp.c
    bench_log.c
    bench_erf.c
    bench_normal.c
)

# Link with benchmark framework and math library
target_link_libraries(math_benchmarks
    PRIVATE
        bench_framework
        finkit_math_static_base
        finkit::platform
        finkit::ds
        m
)

target_compile_definitions(math_benchmarks PRIVATE
    FC_STATIC_BUILD
)

# Include directories
target_include_directories(math_benchmarks
    PRIVATE
        ${CMAKE_SOURCE_DIR}/include
        ${CMAKE_SOURCE_DIR}/modules/platform/include
        ${CMAKE_CURRENT_SOURCE_DIR}
)

# Compiler flags for benchmarks
target_compile_options(math_benchmarks PRIVATE -Wall -Wextra -Wpedantic -ffp-contract=off)

# Installation
install(TARGETS math_benchmarks
    RUNTIME DESTINATION bin
)
