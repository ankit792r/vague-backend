#include "runtime.h"
#include <stdio.h>
#include <stdlib.h>

Runtime *newRuntime(CommandMap *map) {
  Runtime *rt = malloc(sizeof(struct Runtime));
  rt->cmdMap = map;

  return rt;
}

void freeRuntime(Runtime *rt) {
  free(rt->cmdMap);
  free(rt);
}

void startRuntime(Runtime *rt) {
  printf("Runtime started\n");
  while (1) {
  }
}
