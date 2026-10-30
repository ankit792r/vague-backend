#include <stdio.h>
#include <string.h>

void start_loop() {
  int n = 50;
  char inp[n];

  while (1) {
    printf("> ");
    char *in = fgets(inp, n, stdin);
    if (!in) {
      printf("\n");
      return;
    }

    int i = 0;
    while (inp[i] != '\n' && inp[i] != '\0')
      i++;

    if (inp[i] == '\n')
      inp[i] = '\0';

    if (strcmp(inp, "exit") == 0)
      break;

    printf("out: %s\n", inp);
  }
}
