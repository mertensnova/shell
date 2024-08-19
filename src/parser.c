#include "../include/parser.h"
#include "../include/builtin.h"
#include "../include/utils.h"
#include <ctype.h>
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
    micro_exit(n);
    return true;
  };
  if (strncmp("echo", cmd, 4) == 0) {
    s = get_args(input, 4);
    micro_echo(s);
    return true;
  };

  if (strncmp("type", cmd, 4) == 0) {
    s = get_args(input, 4);
    if (!micro_type(s))
      get_path(s);
    return true;
  };

  return false;
};

char *get_args(char *input, int n) {
  size_t size = strlen(input);
  const char *args[100] = {};
  char word[10];
  size_t size1 = strlen(*args);
  char *aa;
  int k = 0;
  int x = 0;
  for (size_t i = 0; i < size; ++i) {
    if (isspace(input[i]) == 0) {
      // printf("%c", input[i]);
      word[k] = input[i];
      k++;
    } else {

      word[k + 1] = '\0';
      args[x] = word;
      k = 0;
      x++;

      printf("%s\n", word);
      memset(word, 0, strlen(word));
    }
  };

  for (size_t i = 0; i < x; ++i) {
    // printf("%s", *args);
  }
  // printf("%s\n", *args);

  return aa;
};
