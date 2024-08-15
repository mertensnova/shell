#include "../include/parser.h"
#include "../include/builtin.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>


bool isbuiltin(char *input) {
  char *s;
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
    s = get_args(input, 4);
    int n = atoi(s);
    nyx_exit(n);
    return true;
  };
  if (strncmp("echo", cmd, 4) == 0) {
    s = get_args(input, 4);
    n_echo(s);
    return true;
  };

  if (strncmp("type", cmd, 4) == 0) {
    s = get_args(input, 4);
        n_type(s);
    return true;
  };

  return false;
};

char *get_args(char *input, int n) {
  size_t size = strlen(input);
  char *args = malloc(sizeof(char) * 20);
  for (size_t i = n; i < size; ++i) {
    args[i - n] = input[i];
  };
  return args;
};
