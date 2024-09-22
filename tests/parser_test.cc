#include "../include/parser.h"
#include <gtest/gtest.h>
#include <vector>
#include <cstring>
#include "../include/builtin.h"
#include "../include/utils.h"

// Test case
TEST(IsBuiltinTest, BasicFunctionality) {

  std::vector<std::string> input = {"type echo", "type exit", "type type"};

  for (int i = 0; i < input.size(); ++i) {
    char *cstr = new char[input[i].length() + 1];
    std::strcpy(cstr, input[i].c_str());
    bool result = isbuiltin((char *)cstr);
    EXPECT_TRUE(result);
  }
}
