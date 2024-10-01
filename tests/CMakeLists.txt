cmake_minimum_required(VERSION 3.29)
include(FetchContent)

FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG        release-1.11.0
)
FetchContent_MakeAvailable(googletest)


# Include the GoogleTest module
include(GoogleTest)

# Create test executables
add_executable(parser_test parser_test.cc)
add_executable(utils_test utils_test.cc)

# Link test executables with GTest and our shell library
target_link_libraries(parser_test PRIVATE GTest::GTest GTest::Main shell_lib)
target_link_libraries(utils_test PRIVATE GTest::GTest GTest::Main shell_lib)

# Discover and add the tests
gtest_discover_tests(parser_test)
gtest_discover_tests(utils_test)

