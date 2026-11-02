#ifndef COMMAND_H
#define COMMAND_H

#define CMD_NODE_SIZE 100

typedef void (*RunFun)();

typedef struct Command {
  char *name;
  char *desc;
  RunFun runfn;
} Command;

typedef struct CommandNode {
  char *key;
  Command cmd;

  struct CommandNode *next;

} CommandNode;

typedef struct CommandMap {
  CommandNode **buckets;
} CommandMap;

Command *GetCommand(CommandMap *map, char *cmd);
void AddCommand(CommandMap *map, Command cmd);

CommandMap *CreateCommandMap();

#endif // !COMMAND_H
