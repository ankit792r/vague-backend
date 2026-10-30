#include "buffer.h"
#include "text.h"
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

buffer_t NewBuffer(char *name, char *path) {

  int columnCount = 20;

  Text_t text;
  text.rowCount = 10;

  text.grid = (char **)malloc(text.rowCount * sizeof(char *));
  for (int i = 0; i < text.rowCount; i++) {
    text.grid[i] = (char *)malloc(columnCount * sizeof(char *));
  }

  buffer_t buff = {.name = name, .path = path, .binary = 0, .text = &text};

  return buff;
}

buffer_t *ScratchBuffer() {
  buffer_t *buff = malloc(sizeof(struct Buffer));
  if (buff == NULL) {
    perror("Failed to allocate buffer");
    return NULL;
  }

  Text_t *text = malloc(sizeof(struct Text));
  if (text == NULL) {
    perror("Failed to allocate text");
    return NULL;
  }
  char tempText[] = "This is scratch buffer\n"
                    "Modified content not saved\n"
                    "Use this for temp buffer\n";

  int rowCount = 0;
  const char *delim = "\n";

  for (int i = 0; tempText[i] != '\0'; i++) {
    if (tempText[i] == '\n') {
      rowCount++;
    }
  }

  char **grid = (char **)malloc(rowCount * sizeof(char *));
  if (grid == NULL) {
    perror("Failed to allocate text grid");
    return NULL;
  }

  char *tok = strtok(tempText, delim);

  int i = 0;
  while (tok != NULL && i < rowCount) {
    grid[i] = (char *)malloc(strlen(tok) * sizeof(char *));
    if (grid[i] == NULL) {
      perror("Failed to allocate row");
      return NULL;
    }

    grid[i] = tok;

    i++;
    tok = strtok(NULL, delim);
  }

  text->rowCount = rowCount;
  text->grid = grid;

  buff->text = text;

  return buff;
}

int LineCount(Text_t *text) { return text->rowCount; }

char *LineAt(Text_t *text, int i) { return text->grid[i]; }
