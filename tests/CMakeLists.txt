#
# tests/CMakeLists.txt - Test configuration for math module
#

# Test executable
add_executable(math_tests
    test_framework.c
    test_main.c
    test_math.c
    test_cumsum.c
    test_cummax.c
    test_exp.c
    test_log.c
    test_erf.c
    test_normal.c
)

# Link with math library
target_link_libraries(math_tests
    PRIVATE
        finkit_math_static_base
        finkit::ds
        m
)

target_compile_definitions(math_tests PRIVATE
    FC_STATIC_BUILD
)

# Include directories
target_include_directories(math_tests
    PRIVATE
        ${CMAKE_SOURCE_DIR}/include
        ${CMAKE_SOURCE_DIR}/modules/platform/include
        ${CMAKE_CURRENT_SOURCE_DIR}
)

# Compiler flags for tests
target_compile_options(math_tests PRIVATE -Wall -Wextra -Wpedantic)

# Run tests with ctest
add_test(
    NAME "math_tests"
    COMMAND math_tests
)

# Installation
install(TARGETS math_tests
    RUNTIME DESTINATION bin
)
