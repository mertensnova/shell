#ifndef PARSER_H
#define PARSER_H


#include <stdbool.h>
bool isbuiltin(char *input);
int get_digit(char *input);

char *get_args(char *input, int n);
#endif
