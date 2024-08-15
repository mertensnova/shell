#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

#define BUFFER_SIZE 1024

char *get_input() {
  char *input = malloc(sizeof(char) * 100);
  char *s = fgets(input, 100, stdin);
  input[strlen(input) - 1] = '\0';

  return input;
};

char *trim_space(char *string) {
  size_t size = sizeof(&string) / sizeof(string[0]);
  for (size_t i = 0; i < size - 1; ++i) {
    if (isspace(string[i])) {
      string[i] = string[i + 1];
    }
  };
  return string;
};
void exe_cmd(char *cmd) {
  if (strlen(cmd) == 0) {
    return;
  };
  char buffer[BUFFER_SIZE];
  int fd[2];
  int fd_pipe = pipe(fd);

  pid_t cpid = fork();
  if (cpid == -1) {
    perror("fork()");
  };

  if (cpid == 0) {
    dup2(fd[1], STDOUT_FILENO);
    execlp(cmd, cmd, NULL);
    printf("%s: command not found", cmd);
    return;
  } else {
    size_t size = read(fd[0], buffer, BUFFER_SIZE);
    buffer[size - 1] = '\0';
    printf("%s", buffer);
  };
  if (kill(cpid, SIGTERM) == 0) {
    return;
  } else {
    perror("Error terminating child process");
  };
};
