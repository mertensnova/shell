#include "../include/parser.h"
#include <cstring>
#include <gtest/gtest.h>
#include <vector>

TEST(IsBuiltinTest, BasicFunctionality) {

  std::vector<std::string> input = {"echo", "exit", "type"};

  for (int i = 0; i < input.size(); ++i) {
    char *cstr = new char[input[i].length() + 1];
    std::strcpy(cstr, input[i].c_str());
    bool result = isbuiltin((char *)cstr);
    EXPECT_TRUE(result);
  }

  std::vector<std::string> input2 = {"pwd", "ls", "nmap"};

  for (int i = 0; i < input2.size(); ++i) {
    char *cstr = new char[input2[i].length() + 1];
    std::strcpy(cstr, input2[i].c_str());
    bool result = isbuiltin((char *)cstr);
    EXPECT_FALSE(result);
  }
}
