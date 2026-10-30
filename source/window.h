#ifndef WINDOW_H
#define WINDOW_H

#include "buffer.h"
#include "cursor.h"

typedef struct Window {
  int id;
  int bufferId;

  marker_t *cursor;
  int topLine;

} window_t;

window_t *NewWindow(int buffId);

void setBuffer(window_t *win, int *buff_id);

#endif // !WINDOW_H
