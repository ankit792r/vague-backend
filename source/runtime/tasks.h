#ifndef TASKS_H
#define TASKS_H

#include <pthread.h>

#define MAX_TASKS 100

typedef struct Task {
  void (*execute)(void *);
  void *data;

  struct Task *task;
} Task;

typedef struct {
  Task *head;
  Task *tail;

  pthread_mutex_t mutex;
  pthread_cond_t condition;

  int shutdown;
} TaskQueue;

// Queue operation
int taskQueuePush(TaskQueue *queue, Task *task);
Task *taskQueuePop(TaskQueue *queue);

// Worker operations
void *worker_loop(void *);

#endif // !TASKS_H
