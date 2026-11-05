#ifndef RUNTIME_H
#define RUNTIME_H

#include "../command.h"

typedef struct Runtime {
  CommandMap *cmdMap;

} Runtime;

// Runtime operations
Runtime *newRuntime(CommandMap *map);
void freeRuntime(Runtime *rt);
void startRuntime(Runtime *rt);

#endif // !RUNTIME_H
