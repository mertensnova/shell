#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

#define BUFFER_SIZE 1024
void path_get(char *cmd);

char *parse_cmd(char *cmd);

int main(void) {

  while (true) {

    printf("\n$ ");
    char input[100];
    char *s = fgets(input, 100, stdin);

    input[strlen(input) - 1] = '\0';

    path_get(input);
  }
  return 0;
};

char *parse_cmd(char *cmd) {
  size_t size = sizeof(&cmd) / sizeof(cmd[0]);
  for (int i = 0; size > 0; ++i) {
    if (cmd[i] == ' ') {
      cmd[i] = '\0';
      return cmd;
    }
  };

  return cmd;
};

void path_get(char *cmd) {
  if (strlen(cmd) == 0) {
    return;
  };

  cmd = parse_cmd(cmd);

  char buffer[BUFFER_SIZE];
  int fd[2];
  int status;
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
