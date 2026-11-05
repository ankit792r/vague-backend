#ifndef QUEUE_H
#define QUEUE_H

#include <pthread.h>

typedef struct Task {
  void (*execute)(void *);
  void *data;

  struct Task *next;
} Task;

typedef struct TaskQueue {
  Task *head;
  Task *tail;

  pthread_mutex_t mutex;
  pthread_cond_t condition;

  int shutdown;

} TaskQueue;

// Queue Operations
int queueInit(TaskQueue *queue);

int queuePush(TaskQueue *queue, Task *task);

Task *queuePop(TaskQueue *queue);

void queueShutdown(TaskQueue *queue);

void queueDestroy(TaskQueue *queue);

#endif // QUEUE_H
