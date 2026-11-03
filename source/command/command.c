#include "command.h"
#include <stdlib.h>
#include <string.h>
static unsigned long hash_function(const char *str) {
  unsigned long hash = 5381;
  int c;

  while ((c = (unsigned char)*str++)) {
    hash = hash * 33 + c;
  }

  return hash;
}

CommandMap *createCommandMap(void) {
  CommandMap *map = malloc(sizeof(*map));

  if (!map)
    return NULL;

  map->capacity = CMD_MAP_CAPACITY;
  map->size = 0;

  map->buckets = calloc(map->capacity, sizeof(*map->buckets));

  if (!map->buckets) {
    free(map);
    return NULL;
  }

  return map;
}

int addCommand(CommandMap *map, Command cmd) {
  if (!map || !cmd.name || !cmd.runFun) {
    return -1;
  }

  unsigned long index = hash_function(cmd.name) % map->capacity;

  CommandNode *head = map->buckets[index];

  // Check for existing command
  while (head != NULL) {
    if (strcmp(head->cmd.name, cmd.name) == 0) {
      head->cmd = cmd;
      return 0;
    }

    head = head->next;
  }

  // Allocate new node
  CommandNode *node = malloc(sizeof(*node));

  if (!node) {
    return -1;
  }

  node->cmd = cmd;
  node->next = map->buckets[index];

  map->buckets[index] = node;
  map->size++;

  return 0;
}

Command *getCommand(CommandMap *map, char *cmd) {
  if (!map || !cmd)
    return NULL;

  unsigned long index = hash_function(cmd) % map->capacity;
  CommandNode *head = map->buckets[index];

  while (head != NULL) {
    if (strcmp(head->cmd.name, cmd) == 0)
      return &head->cmd;

    head = head->next;
  }

  return NULL;
}

void destroyCommandMap(CommandMap *map) {
  if (!map)
    return;

  for (size_t i = 0; i < map->capacity; i++) {
    CommandNode *node = map->buckets[i];
    while (node != NULL) {
      CommandNode *next = node->next;
      // commented below code because it cassing use after free
      // free(next);
      // node = next;
    }
  }

  free(map->buckets);
  free(map);
}
