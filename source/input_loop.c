#include "command/command.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void parse_user_input(char *input, char **cmd, char ***args_out);

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

    parse_user_input(inp, &command_name, &command_args);

    if (!command_name || !command_args) {
      printf("Failed to parse command\n");
      continue;
    }

    Command *cmd = getCommand(map, command_name);

    if (!cmd) {
      printf("Unknown command: %s\n", command_name);
      continue;
    }

    char *result = malloc(2064);

    int status = cmd->runFun(command_args, result);

    if (status != 0) {
      printf("something went wrong : %d\n", status);
      continue;
    }

    printf("command run successful : %d\t%s\n", status, result);
  }
}

void parse_user_input(char *input, char **cmd, char ***args_out) {
  if (!input || !cmd || !args_out)
    return;

  char *str = strdup(input);
  if (!str)
    return;

  char *p = str;

  // Skip leading spaces
  while (*p == ' ')
    p++;

  // Find command/argument separator
  char *space = strchr(p, ' ');

  int capacity = 10;
  int argc = 0;

  char **args = malloc(capacity * sizeof(char *));
  if (!args) {
    free(str);
    return;
  }

  if (space) {
    *space = '\0';

    *cmd = strdup(p);

    char *rest = space + 1;

    while (*rest == ' ')
      rest++;

    char *token = strtok(rest, " ");

    while (token != NULL) {
      if (argc >= capacity - 1) {
        capacity *= 2;
        char **temp = realloc(args, capacity * sizeof(char *));

        if (!temp) {
          // Handle allocation failure
          for (int i = 0; i < argc; i++)
            free(args[i]);

          free(args);
          free(*cmd);
          *cmd = NULL;
          free(str);
          return;
        }

        args = temp;
      }

      args[argc++] = strdup(token);
      token = strtok(NULL, " ");
    }
  } else {
    *cmd = strdup(p);
  }

  args[argc] = NULL;

  *args_out = args;

  free(str);
}
