#include "command.h"
#include <stdio.h>

void registerCommands(CommandMap *map) {
  RegisterDefaultCommands(map);
  RegisterBufferCommands(map);
}

static int cmdAllCommands(char **args, char *res) {
  res = "we will list all command here";
  return 0;
}

static int cmdDescribeCommand(char **args, char *res) {
  printf("received params: %s\n", *args);
  return 0;
}

void RegisterDefaultCommands(CommandMap *map) {
  addCommand(map, (Command){.name = "all-cmd",
                            .desc = "list all commands",
                            .runFun = cmdAllCommands});

  addCommand(map, (Command){.name = "desc-cmd",
                            .desc = "describe specific command",
                            .runFun = cmdDescribeCommand});
}

void RegisterBufferCommands(CommandMap *map) {}
