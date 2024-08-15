#ifndef BUILTIN_H
#define BUILTIN_H

#include <stdbool.h>

extern const char *builtin[];

void n_type(char *arg);
void nyx_exit(int status);
void n_echo(char *messeage);
#endif
