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
    char **s = get_args(input);
    int n = atoi((char *)s);
    micro_exit(n);
    return true;
  };
  if (strncmp("echo", cmd, 4) == 0) {
    char **s = get_args(input);
    micro_echo(s);
    return true;
  };

  if (strncmp("type", cmd, 4) == 0) {
    char **s = get_args(input);
    if (micro_type(s[1]))
      printf("%s is a shell builtin", s[1]);
    else
      get_path(s[1]);

    return true;
  };

  return false;
};
char **get_args(char *input) {
  size_t size = strlen(input);
  char **args;
  args = (char **)malloc(100 * sizeof(char *));
  int k = 0;
  int j = 0;
  char *word = malloc(sizeof(char) * (size + 1)); // +1 for null terminator
  for (size_t i = 0; i < size; i++) {
    if (input[i] == '"') {
      continue;
    }
    if (isspace(input[i]) == 0) {
      word[k++] = input[i];
    } else {
      word[k] = '\0';
      args[j] = (char *)malloc(sizeof(char) * k + 1);
      strcpy(args[j], word);
      k = 0;
      j++;
    };
  };

  if (k > 0) {
    word[k] = '\0';
    args[j] = (char *)malloc(sizeof(char) * (k + 1));
    if (args[j] == NULL) {
      perror("Failed to allocate memory");
      exit(EXIT_FAILURE);
    }
    strcpy(args[j], word);
    j++;
  }

  free(word);

  return args;
};
