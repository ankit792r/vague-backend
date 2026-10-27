#include "buffer.h"
#include "text.h"
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>

Buffer_t NewBuffer(char *name, char *path) {

  int columnCount = 20;

  Text_t text;
  text.rowCount = 10;

  text.grid = (char **)malloc(text.rowCount * sizeof(char *));
  for (int i = 0; i < text.rowCount; i++) {
    text.grid[i] = (char *)malloc(columnCount * sizeof(char *));
  }

  Buffer_t buff = {.name = name, .path = path, .binary = 0, .text = &text};

  return buff;
}

Buffer_t ScratchBuffer() {
  Text_t *text = malloc(sizeof(*text));

  if (text == NULL)
    return (Buffer_t){0};

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

  char *tok = strtok(tempText, delim);

  int i = 0;
  while (tok != NULL && i < rowCount) {
    text->grid[i++] = tok;
    tok = strtok(NULL, delim);
  }

  text->rowCount = i;

  Buffer_t buff = {.name = "scratch", .binary = 0, .text = text};

  return buff;
}

int LineCount(Text_t *text) { return text->rowCount; }

char *LineAt(Text_t *text, int i) { return text->grid[i]; }
