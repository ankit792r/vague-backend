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

typedef struct {
  CommandNode *bucket[CMD_NODE_SIZE];
} CommandMap;

void GetCommand(char *cmd);
void AddCommand(Command cmd);

#endif // !COMMAND_H
