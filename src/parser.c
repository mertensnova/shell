#include "../include/parser.h"
#include "../include/builtin.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

bool isbuiltin(char *input) {
  size_t size = strlen(input);
  char cmd[100];

  for (size_t i = 0; i < size; ++i) {
    if (input[i] != '\t' && input[i] != ' ') {
      cmd[i] = input[i];
    } else {
      break;
    };
  };

  if (strncmp("exit", cmd, 4) == 0) {
    int n = get_digit(input);
    nyx_exit(n);
    return true;
  };
  return false;
};

int get_digit(char *input) {
  size_t size = strlen(input);
  char digit[10];

  for (size_t i = 4; i < size; ++i) {
    digit[i - 4] = input[i];
  };

  int num = atoi(digit);
  return num;
};
