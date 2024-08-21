#include "../include/parser.h"
#include "../include/utils.h"
#include <stdbool.h>
#include <stdio.h>
#include <sys/wait.h>
#include <unistd.h>

int main(void) {
  char *input;
  do {
    printf("\n\n$ ");
    input = get_input();
    //   if (isbuiltin(input)) {
    //   continue;
    //};
    get_args(input, 3);
    // exe_cmd(input);
  } while (true);
  return 0;
};
