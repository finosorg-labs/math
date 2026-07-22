cmake_minimum_required(VERSION 3.16)
project(finkit-math VERSION 1.0.0 LANGUAGES C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

# Build options (must be before add_subdirectory calls)
option(FC_BUILD_TESTS "Build tests" ON)
option(FC_BUILD_BENCHMARKS "Build benchmarks" ON)

# Include SIMD detection and compiler flags
include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/DetectSIMD.cmake)
include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/CompilerFlags.cmake)

# Sanitizer configuration
if(FC_ENABLE_SANITIZERS)
    if(EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/cmake/sanitizers.cmake)
        include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/sanitizers.cmake)
    endif()
endif()

# Generate cgo_source.go with detected SIMD flags
if(FC_SIMD_LEVEL STREQUAL "AVX512")
    set(FC_SIMD_CFLAGS "${FC_AVX512_FLAGS}")
elseif(FC_SIMD_LEVEL STREQUAL "AVX2")
    set(FC_SIMD_CFLAGS "${FC_AVX2_FLAGS}")
elseif(FC_SIMD_LEVEL STREQUAL "SSE42")
    set(FC_SIMD_CFLAGS "${FC_SSE42_FLAGS}")
else()
    set(FC_SIMD_CFLAGS "")
endif()
configure_file(
    ${CMAKE_CURRENT_SOURCE_DIR}/cgo_source.go.in
    ${CMAKE_CURRENT_SOURCE_DIR}/cgo_source.go
    @ONLY
)


# Build math-c subdirectory (creates math_obj target)
add_subdirectory(math-c)

# Create base static library (without platform)
add_library(finkit_math_static_base STATIC $<TARGET_OBJECTS:math_obj>)
target_include_directories(finkit_math_static_base
    PUBLIC
        $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
        $<INSTALL_INTERFACE:include>
)

# Determine math build directory
if(CMAKE_SYSTEM_NAME STREQUAL "Darwin")
    if(CMAKE_SYSTEM_PROCESSOR MATCHES "arm64|aarch64")
        set(MATH_BUILD_DIR "darwin_arm64")
    else()
        set(MATH_BUILD_DIR "darwin_x86_64")
    endif()
elseif(CMAKE_SYSTEM_NAME STREQUAL "Linux")
    set(MATH_BUILD_DIR "linux_amd64")
elseif(CMAKE_SYSTEM_NAME STREQUAL "Windows")
    set(MATH_BUILD_DIR "windows_amd64")
endif()

# Merge static libraries with deduplication
# Use build-specific path to avoid conflicts between different build types (sanitizers)
set(MERGED_LIB_PATH ${CMAKE_CURRENT_BINARY_DIR}/libfinkit_math_static.a)
# Fixed path for Go cgo bindings
set(FINAL_LIB_PATH ${CMAKE_CURRENT_SOURCE_DIR}/build/${MATH_BUILD_DIR}/libfinkit_math_static.a)

# Determine merge tool based on platform
if(CMAKE_SYSTEM_NAME STREQUAL "Windows")
    set(MERGE_TOOL "lib.exe")
elseif(CMAKE_SYSTEM_NAME STREQUAL "Darwin")
    set(MERGE_TOOL "libtool")
else()
    set(MERGE_TOOL "${CMAKE_AR}")
endif()

# Unified merge command for all platforms
set(MERGE_SCRIPT ${CMAKE_CURRENT_SOURCE_DIR}/scripts/merge_libs.sh)

add_custom_command(
    OUTPUT ${MERGED_LIB_PATH}
    COMMAND ${CMAKE_COMMAND} -E make_directory ${CMAKE_CURRENT_SOURCE_DIR}/build/${MATH_BUILD_DIR}
    COMMAND bash ${MERGE_SCRIPT}
        ${MERGED_LIB_PATH}
        $<TARGET_FILE:finkit_math_static_base>
        ${MERGE_TOOL}
        --mode=smart
        --project-root=${CMAKE_CURRENT_SOURCE_DIR}
        --build-subdir=${MATH_BUILD_DIR}
    COMMAND ${CMAKE_COMMAND} -E copy_if_different ${MERGED_LIB_PATH} ${FINAL_LIB_PATH}
    DEPENDS finkit_math_static_base ${MERGE_SCRIPT}
    COMMENT "Merging static libraries with smart deduplication"
    VERBATIM
)

# Create custom target to trigger the merge
add_custom_target(finkit_math_static ALL
    DEPENDS ${MERGED_LIB_PATH}
    COMMAND ${CMAKE_COMMAND} -E echo "Merged library created at: ${MERGED_LIB_PATH}"
    COMMAND ${CMAKE_COMMAND} -E echo "Final library for Go: ${FINAL_LIB_PATH}"
)

# Create interface library for linking
add_library(finkit_math_static_interface INTERFACE)
target_include_directories(finkit_math_static_interface
    INTERFACE
      $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
        $<INSTALL_INTERFACE:include>
)
target_link_libraries(finkit_math_static_interface
    INTERFACE
        ${MERGED_LIB_PATH}
        m
)

# Create alias
add_library(finkit::math ALIAS finkit_math_static_interface)

# Install merged static library
install(FILES ${MERGED_LIB_PATH}
    DESTINATION lib
)

install(DIRECTORY include/
    DESTINATION include/math
    FILES_MATCHING PATTERN "*.h"
)

# Install interface library target
install(TARGETS finkit_math_static_interface
    EXPORT finkit-math-targets
)

install(EXPORT finkit-math-targets
    FILE finkit-mathTargets.cmake
    NAMESPACE finkit::
    DESTINATION lib/cmake/finkit-math
)

# Config files
include(CMakePackageConfigHelpers)
write_basic_package_version_file(
    "${CMAKE_CURRENT_BINARY_DIR}/finkit-mathConfigVersion.cmake"
    VERSION ${PROJECT_VERSION}
    COMPATIBILITY AnyNewerVersion
)

configure_package_config_file(
    "${CMAKE_CURRENT_SOURCE_DIR}/cmake/finkit-mathConfig.cmake.in"
    "${CMAKE_CURRENT_BINARY_DIR}/finkit-mathConfig.cmake"
    INSTALL_DESTINATION lib/cmake/finkit-math
)

install(FILES
    "${CMAKE_CURRENT_BINARY_DIR}/finkit-mathConfig.cmake"
    "${CMAKE_CURRENT_BINARY_DIR}/finkit-mathConfigVersion.cmake"
    DESTINATION lib/cmake/finkit-math
)


# Tests
if(FC_BUILD_TESTS AND EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/tests)
    enable_testing()
    add_subdirectory(tests)
endif()

# Benchmarks
if(FC_BUILD_BENCHMARKS AND EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/benchmarks)
    add_subdirectory(benchmarks)
endif()

# Print configuration
message(STATUS "")
message(STATUS "============================================================")
message(STATUS " finkit-math ${PROJECT_VERSION}")
message(STATUS "============================================================")
message(STATUS "  C Compiler: ${CMAKE_C_COMPILER_ID} ${CMAKE_C_COMPILER_VERSION}")
message(STATUS "  Build Type: ${CMAKE_BUILD_TYPE}")
message(STATUS "  System: ${CMAKE_SYSTEM_NAME}")
message(STATUS "============================================================")
message(STATUS "")
