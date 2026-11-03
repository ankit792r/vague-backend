#include "command/command.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void parse_user_input(char *input, char *cmd, char **args_out);

void start_loop(CommandMap *map) {
  int n = 50;
  char inp[n];

  while (1) {
    printf("> ");
    char *in = fgets(inp, n, stdin);
    if (!in) {
      printf("\n");
      return;
    }

    int i = 0;
    while (inp[i] != '\n' && inp[i] != '\0')
      i++;

    if (inp[i] == '\n')
      inp[i] = '\0';

    if (strcmp(inp, "exit") == 0)
      break;

    char *command_name = NULL;
    char **command_args = NULL;

    parse_user_input(inp, command_name, command_args);

    Command *cmd = getCommand(map, command_name);

    char *result = malloc(2064);

    int status = cmd->runFun(command_args, result);

    if (status != 0) {
      printf("something went wrong : %d\n", status);
      continue;
    }

    printf("command run successful : %d\t%s\n ", status, result);
  }
}

void parse_user_input(char *input, char *cmd, char **args_out) {
  if (!input || !cmd || !args_out)
    return;

  // Create a modifiable copy of the input string
  char *str = strdup(input);
  if (!str)
    return;

  // Skip leading spaces
  char *p = str;
  while (*p == ' ')
    p++;

  // Find the first space separating the command and arguments
  char *space = strchr(p, ' ');
  if (space) {
    *space = '\0';
    cmd = strdup(p);

    // Move past the space and any extra whitespace for arguments
    char *rest = space + 1;
    while (*rest == ' ')
      rest++;

    // Tokenize the remaining arguments into a dynamic char** array
    int capacity = 10;
    int argc = 0;
    char **args = malloc(capacity * sizeof(char *));

    char *token = strtok(rest, " ");
    while (token != NULL) {
      if (argc >= capacity - 1) {
        capacity *= 2;
        args = realloc(args, capacity * sizeof(char *));
      }
      args[argc++] = strdup(token);
      token = strtok(NULL, " ");
    }
    args[argc] = NULL; // NULL-terminate the array (essential for execvp, etc.)
    *args_out = *args;
  } else {
    // No arguments provided, only a command
    cmd = strdup(p);
    char **args = malloc(sizeof(char *));
    args[0] = NULL;
    *args_out = *args;
  }

  free(str);
}
