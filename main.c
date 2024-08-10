#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

#define MAX_PATH 1024
void path_get(char *cmd);

int main(void) {
  printf("\n$ ");

  char input[100];
  char *s = fgets(input, 100, stdin);

  input[strlen(input) - 1] = '\0';

  path_get(input);
  return 0;
};

void path_get(char *cmd) {

  char buffer[1000];
  int fd[2];
  int status;
  int fd_pipe = pipe(fd);

  pid_t cpid = fork();
  if (cpid == -1) {
    perror("fork()");
  };

  if (cpid == 0) {
    //dup2(fd[1], STDOUT_FILENO);
    execlp(cmd, cmd, NULL);
   // printf("%s command not found", cmd);
  }};
