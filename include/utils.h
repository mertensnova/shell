#ifndef UTILS_H
#define UTILS_H

#ifdef __cplusplus
extern "C" {
#endif

char *get_input();
void get_path(char *cmd);
void exe_cmd(char *cmd);
char *trim_space(char *string);
void exec(char *string);

#ifdef __cplusplus
}
#endif
#endif
