#include "buffer.h"
#include "command.h"
#include "input_loop.h"
#include "window.h"
#include <stdio.h>
#include <stdlib.h>

void echoFn() { printf("Hello there\n"); }
void versionFn() { printf("version v0.0.0\n"); }

int main(void) {
  buffer_t *scratch = ScratchBuffer();
  window_t *win = NewWindow(3);

  CommandMap *map = CreateCommandMap();
  Command echoCmd = {
      .name = "echo", .desc = "Echo hello message", .runfn = *echoFn};
  Command versionCmd = {
      .name = "version", .desc = "Echo current version", .runfn = *versionFn};

  AddCommand(map, echoCmd);
  AddCommand(map, versionCmd);

  start_loop(map);

  free(scratch->text);
  free(scratch);

  free(win->cursor);
  free(win);
  return 0;
}
