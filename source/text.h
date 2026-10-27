#ifndef TEXT_H
#define TEXT_H

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

#endif // !TEXT_H
