#ifndef RUNTIME_H
#define RUNTIME_H

#include "command/command.h"

typedef struct Runtime {
  CommandMap *cmdMap;

} Runtime;

Runtime *newRuntime(CommandMap *map);
void freeRuntime(Runtime *rt);
void startRuntime(Runtime *rt);

#endif // !RUNTIME_H
