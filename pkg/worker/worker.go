package worker

import (
	"os"
	"strconv"
	"time"

	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/entity"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/queue"
)

type Worker struct {
	repo  MetricRepo
	queue *queue.Queue
}

type MetricRepo interface {
	Sub(entity.Metric)
}

func NewWorker(repo MetricRepo, queue *queue.Queue) *Worker {
	return &Worker{
		repo:  repo,
		queue: queue,
	}
}

func (w *Worker) Run() {
	timer, err := strconv.Atoi(os.Getenv("TIME_TICKER_IN_SECONDS"))
	if err != nil {
		timer = 5 // default time if something went wrong
	}

	duration, err := strconv.Atoi(os.Getenv("METRICS_DURATION_IN_SECONDS"))
	if err != nil {
		timer = 3600 // default time if something went wrong
	}

	metricsExpiration := time.Duration(duration * int(time.Second))

	for {
		// sleeps before checking metrics that have already expired
		time.Sleep(time.Duration(timer) * time.Second)

		for {
			metric, err := w.queue.GetNext()
			if err != nil {
				break
			}

			now := time.Now()
			if now.After(metric.CreatedAt.Add(metricsExpiration)) {
				w.repo.Sub(metric)
				w.queue.Dequeue()
			} else {
				break
			}
		}
	}
}
