#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void path_get(char *cmd);
int main(void) {
  while (true) {
    printf("\n$ ");
    fflush(stdout);

    char input[100];
    char *s = fgets(input, 100, stdin);

    char *args[] = {input, NULL};
    input[strlen(input) - 1] = '\0';

    path_get(input);
    /*
        if (execvp("/usr/bin", args) == -1) {
          printf("%s: command not found\n", input);
          continue;
        };
        */
  };
  return 0;
};
/*
  pid_t pid = fork();
  if (pid == 0) {
    char *args[] = {input, NULL};
    if (execvp("/usr/bin/env", args) == -1)
         printf("%s: command not found", input);
  };
  */

void path_get(char *cmd) {
  pid_t pid = fork();
  if (pid == 0) {
    char *args[] = {"which",cmd, NULL};
    if (execvp("/usr/bin", args) == -1)
      printf("\npath of %s not found", cmd);
  };
};
