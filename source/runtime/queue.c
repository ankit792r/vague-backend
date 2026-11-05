#include "queue.h"

int queueInit(TaskQueue *queue) {
  queue->head = NULL;
  queue->tail = NULL;
  queue->shutdown = 0;

  if (pthread_mutex_init(&queue->mutex, NULL) != 0) {
    return -1;
  }

  if (pthread_cond_init(&queue->condition, NULL) != 0) {
    pthread_mutex_destroy(&queue->mutex);
    return -1;
  }

  return 0;
}
