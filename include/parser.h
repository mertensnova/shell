#ifndef PARSER_H
#define PARSER_H


#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

bool isbuiltin(char* input);

#ifdef __cplusplus
}
#endif
int get_digit(char *input);

char **get_args(char *input);
#endif
