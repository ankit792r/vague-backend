#include "window.h"
#include "cursor.h"
#include <stdlib.h>

window_t *NewWindow(int buffId) {
  window_t *win = malloc(sizeof(struct Window));
  win->id = 0;
  win->bufferId = buffId;

  win->cursor = malloc(sizeof(struct Marker));
  win->cursor->offset = 0;

  return win;
}
