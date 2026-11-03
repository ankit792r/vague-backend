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
  Command cmd;
  struct CommandNode *next;
} CommandNode;

typedef struct CommandMap {
  CommandNode **buckets;
} CommandMap;

CommandMap *CreateCommandMap();
Command *GetCommand(CommandMap *map, char *cmd);
void AddCommand(CommandMap *map, Command cmd);
void RegisterCommands(CommandMap *map);

#endif // !COMMAND_H
