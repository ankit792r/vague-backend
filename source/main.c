#include "buffer.h"
#include <stdio.h>

int main(void) {
  Buffer_t scratch = ScratchBuffer();

  printf("lines %d", LineCount(scratch.text));

  for (int i = 0; i < LineCount(scratch.text); i++) {
    printf("%s", LineAt(scratch.text, i));
  }

  return 0;
}
