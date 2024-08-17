#ifndef BUILTIN_H
#define BUILTIN_H

#include <stdbool.h>

extern const char *builtin[];

bool micro_type(char *arg);
void micro_exit(int status);
void micro_echo(char *messeage);
#endif
