#include "buffer.h"
#include <stdio.h>

int main(void) {
  Buffer_t *scratch = ScratchBuffer();

  printf("lines %d\n", LineCount(scratch->text));
  printf("line = %s", LineAt(scratch->text, 0));

  // for (int i = 0; i < LineCount(scratch->text); i++) {
  //   printf("%s", LineAt(scratch->text, i));
  // }

  return 0;
}
