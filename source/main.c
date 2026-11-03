#include "buffer.h"
#include "command/command.h"
#include "input_loop.h"
#include "window.h"
#include <stdlib.h>

int main(void) {
  buffer_t *scratch = ScratchBuffer();
  window_t *win = NewWindow(3);

  CommandMap *map = createCommandMap();

  registerCommands(map);

  start_loop(map);

  free(scratch->text);
  free(scratch);

  free(win->cursor);
  free(win);
  return 0;
}
