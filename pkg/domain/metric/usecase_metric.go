package metric

import (
	"time"

	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/entity"
	"github.com/lndaquino/true-tickets-sum-metrics/pkg/queue"
)

type MetricUsecase struct {
	repo  MetricRepo
	queue *queue.Queue
}

type MetricRepo interface {
	Sum(entity.Metric)
	GetValue(string) int
}

func NewMetricUsecase(repo MetricRepo, queue *queue.Queue) *MetricUsecase {
	return &MetricUsecase{
		repo:  repo,
		queue: queue,
	}
}

func (usecase *MetricUsecase) Get(key string) int {
	return usecase.repo.GetValue(key)
}

func (usecase *MetricUsecase) Create(metric entity.Metric) {
	metric.CreatedAt = time.Now()
	usecase.queue.Enqueue(metric)
	usecase.repo.Sum(metric)
}
