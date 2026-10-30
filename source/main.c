#include "buffer.h"
#include "input_loop.h"
#include "window.h"
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  buffer_t *scratch = ScratchBuffer();
  window_t *win = NewWindow(3);

  printf("window %d\n", win->id);
  printf("buffer %d\n", win->bufferId);

  start_loop();

  free(scratch->text);
  free(scratch);

  free(win->cursor);
  free(win);
  return 0;
}
