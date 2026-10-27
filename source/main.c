#include <stdlib.h>

typedef int Offset_t;

typedef struct {
  int row;
  int column;
} TextPoint;

typedef struct {
  int rowCount;
  char **grid;

  Offset_t starts[];
} Text_t;

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

int main(void) {
  int columnCount = 20;

  Text_t text;
  text.rowCount = 10;

  text.grid = (char **)malloc(text.rowCount * sizeof(char *));
  for (int i = 0; i < text.rowCount; i++) {
    text.grid[i] = (char *)malloc(columnCount * sizeof(char *));
  }

  Buffer_t buffer;
  buffer.text = &text;

  return 0;
}
