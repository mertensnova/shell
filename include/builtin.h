#ifndef BUILTIN_H
#define BUILTIN_H

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif
extern const char *builtin[];

bool micro_type(char *arg);
void micro_exit(int status);
void micro_echo(char **messeage);

#ifdef __cplusplus
}
#endif
#endif
