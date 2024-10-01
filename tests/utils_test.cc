#include "../include/utils.h"
#include <cstring>
#include <gtest/gtest.h>

/*
TEST(IsBuiltinTest, BasicFunctionality) {
  const char *builtins[] = {"echo", "type", "exist", nullptr};

  for (int i = 0; builtins[i] != nullptr; ++i) {
    bool result = isbuiltin(builtins[i]);
    EXPECT_TRUE(result);
  }
}
*/

TEST(TestGetInput, TestInput) { 
    const char *input = "pwd";

    char *output = get_input((char *)input);
}
