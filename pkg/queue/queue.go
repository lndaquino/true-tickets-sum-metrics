package queue

import (
	"errors"
	"log"

	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/entity"
)

type Queue struct {
	queue []entity.Metric
}

func NewQueue() *Queue {
	return &Queue{queue: nil}
}

func (q *Queue) Enqueue(m entity.Metric) {
	log.Printf("Enqueueing metric ==> %v", m)
	q.queue = append(q.queue, m)
}

func (q *Queue) GetNext() (entity.Metric, error) {
	if len(q.queue) > 0 {
		return q.queue[0], nil
	}
	return entity.Metric{}, errors.New("Empty queue")

}

func (q *Queue) Dequeue() {
	log.Printf("Dequeueing...")
	q.queue = q.queue[1:]
}
