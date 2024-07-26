#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define MAX_PATH 1024
void path_get(char *cmd);
int redirect();
int main(void) {
  while (true) {
    printf("\n$ ");
    fflush(stdout);

    char input[100];
    char *s = fgets(input, 100, stdin);

    char *args[] = {input, NULL};
    input[strlen(input) - 1] = '\0';

    path_get(input);
    
  };
  return 0;
};

int redirect() { return 0; };

void path_get(char *cmd) {

  int fd[2];
  pipe(fd);
  if (fork() == 0) {
    dup2(fd[1], STDOUT_FILENO);
    execlp(cmd, cmd, NULL);
    printf("%s: command not found\n", cmd);
  } else {
    char buffer[10000];
    ssize_t size = read(fd[0], buffer, 10000);
    if ((size > 0) && (size < sizeof(buffer))) {
      buffer[size] = '\0';
      printf("%s\n", buffer);
    }
  }
  redirect();
};
