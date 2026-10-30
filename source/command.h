#ifndef COMMAND_H
#define COMMAND_H

#define CMD_NODE_SIZE 100

typedef struct Command {
  char *name;
  char *desc;
  void (*run)();
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
