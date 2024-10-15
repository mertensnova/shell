#include "../include/builtin.h"
#include "../include/utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const char *builtin[] = {"echo", "type", "exit"};

void micro_exit(int status) { exit(status); }
void micro_echo(char **messeage) {
  size_t size = sizeof(messeage);
  for (size_t i = 1; i < size; ++i) {

    if (messeage[i] == NULL) {
      continue;
    };

    printf("%s", messeage[i]);
    printf(" ");
  }
};

bool micro_type(char *arg) {
  size_t size = sizeof(builtin) / sizeof(builtin[0]);

  for (size_t i = 0; i < size; ++i) {
    size_t arg_size = sizeof(&arg) / sizeof(arg[0]);
    size_t builtin_size = sizeof(&builtin[i]) / sizeof(builtin[i][0]);
    if (arg_size == builtin_size) {
      if (strncmp(trim_space(arg), builtin[i], arg_size) == 0) {
        printf("%s is a shell builtin", arg);
        return true;
      }
    }
  };

  return false;
}
