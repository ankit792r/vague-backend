#include "command.h"
#include <stdio.h>
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

    Command *cmd = GetCommand(map, inp);
    if (cmd == NULL) {
      printf("Command not found : `%s`\n", inp);
      continue;
    }

    cmd->runfn();

    // printf("out: %s\n", inp);
  }
}
