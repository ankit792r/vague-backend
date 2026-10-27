#ifndef BUFFER_H
#define BUFFER_H

#include "text.h"

typedef enum { LF, CRLF } LineEnding_t;

typedef struct {
  char *name;
  char *path;

  Text_t *text;
  // add undo tree

  int noEOL;
  LineEnding_t lineEnding;

  int binary;
  int readOnly;
} Buffer_t;

Buffer_t NewBuffer(char *name, char *path);

Buffer_t ScratchBuffer();

int LineCount(Text_t *text);

char *LineAt(Text_t *text, int i);
#endif // !BUFFER_H
