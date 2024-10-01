#include <ctype.h>
#include <fcntl.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

#define BUFFER_SIZE 2024

char *get_input() {
  char *input = malloc(sizeof(char) * 100);
  char *s = fgets(input, 100, stdin);
  input[strlen(input) - 1] = '\0';

  return input;
};

void get_path(char *cmd) {
  int fd[2];
  int fd_pipe = pipe(fd);

  char buffer[BUFFER_SIZE];
  char path[256] = "/usr/bin/";
  strcat(path, cmd);
  pid_t pid = fork();

  if (pid == 0) {
    dup2(fd[1], STDOUT_FILENO);
    char *args[] = {"which", cmd, NULL};
    int nfd[3];
    int pip = pipe(nfd);
    pid_t npid = fork();
    if (npid == 0) {
      dup2(nfd[2], STDERR_FILENO);
      execvp("which", args);
    } else {
      printf("%s: not found", cmd);
      return;
    };

  } else {
    size_t size = read(fd[0], buffer, BUFFER_SIZE);
    buffer[size - 1] = '\0';
    printf("%s is %s", cmd, buffer);
    return;
  };
  return;
};

char *trim_space(char *string) {
  size_t size = sizeof(&string) / sizeof(string[0]);
  int i = 0;
  while (isspace(string[i]) && i < size) {
    char a, tmp;
    tmp = string[i];
    a = string[i + 1];
    string[i] = a;
    string[i + 1] = tmp;
    i++;
  }
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

