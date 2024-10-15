#ifndef PARSER_H
#define PARSER_H

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

bool isbuiltin(const char *input);
int get_digit(char *input);
char *get_args(char *input);

#ifdef __cplusplus
}
#endif
#endif
