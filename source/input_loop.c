#include "command/command.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

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

    // here split to <cmd> <args to the cmd>
    strtok(inp, " ");

    Command *cmd = getCommand(map, inp);
    if (cmd == NULL) {
      printf("Command not found : `%s`\n", inp);
      continue;
    }

    char *t = "-----------------";
    char *res = malloc(2064);

    int status = cmd->runFun(&t, res);

    if (status != 0) {
      printf("something went wrong : %d\n", status);
      continue;
    }

    printf("command run successful : %d\t%s\n ", status, res);

    // printf("out: %s\n", inp);
  }
}
