#ifndef COMMAND_H
#define COMMAND_H

#include <stddef.h>
#define CMD_MAP_CAPACITY 2048

typedef int (*RunFun)(char **args, char *res);

typedef struct Command {
  const char *name;
  const char *desc;
  RunFun runFun;
} Command;

typedef struct CommandNode {
  Command cmd;
  struct CommandNode *next;
} CommandNode;

typedef struct CommandMap {
  CommandNode **buckets;
  size_t capacity;
  size_t size;
} CommandMap;

CommandMap *createCommandMap(void);

Command *getCommand(CommandMap *map, char *cmd);

int addCommand(CommandMap *map, Command cmd);

void registerCommands(CommandMap *map);

void destroyCommandMap(CommandMap *map);


// command register helper for other modules
void RegisterDefaultCommands(CommandMap *map);
void RegisterBufferCommands(CommandMap *map);
#endif // !COMMAND_H
