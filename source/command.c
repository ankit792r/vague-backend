#include "command.h"
#include <stdlib.h>
#include <string.h>

unsigned long hash_function(const char *str) {
  unsigned long hash = 5381;
  int c;
  while ((c = *str++)) {
    hash = ((hash << 5) + hash) + c; // hash * 33 + c
  }
  return hash % CMD_NODE_SIZE;
}

CommandMap *CreateCommandMap() {
  CommandMap *map = (CommandMap *)malloc(sizeof(CommandMap));
  map->buckets = (CommandNode **)calloc(CMD_NODE_SIZE, sizeof(CommandNode *));

  return map;
}

void AddCommand(CommandMap *map, Command cmd) {
  unsigned long idx = hash_function(cmd.name);
  CommandNode *head = map->buckets[idx];

  while (head != NULL) {
    if (strcmp(head->key, cmd.name) == 0) {
      head->cmd = cmd;
      return;
    }
    head = head->next;
  }

  CommandNode *newNode = malloc(sizeof(CommandNode));
  newNode->key = strdup(cmd.name);
  newNode->cmd = cmd;
  newNode->next = map->buckets[idx];

  map->buckets[idx] = newNode;
}

Command *GetCommand(CommandMap *map, char *cmd) {
  unsigned long index = hash_function(cmd);
  CommandNode *head = map->buckets[index];

  Command command;

  while (head != NULL) {
    if (strcmp(head->cmd.name, cmd) == 0) {
      return &head->cmd;
    }
    head = head->next;
  }

  return NULL;
}

void HandleCommand(char *cmd) {}

void RegisterCommands(CommandMap *map) {
  `

}