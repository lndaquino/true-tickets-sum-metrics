package datastore

import (
	"log"
	"sync"

	"github.com/lndaquino/true-tickets-sum-metrics/pkg/domain/entity"
)

type InMemoryRepo struct {
	mutex *sync.Mutex
	m     map[string]int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		m:     map[string]int{},
		mutex: &sync.Mutex{},
	}
}

func (repo *InMemoryRepo) Sum(metric entity.Metric) {
	log.Printf("Adding metric ==> %v", metric)
	repo.mutex.Lock()
	_, found := repo.m[metric.Key]

	if found {
		repo.m[metric.Key] += metric.Value
	} else {
		repo.m[metric.Key] = metric.Value
	}
	repo.mutex.Unlock()
}

func (repo *InMemoryRepo) Sub(metric entity.Metric) {
	log.Printf("Subtracting metric ==> %v", metric)
	repo.mutex.Lock()
	_, found := repo.m[metric.Key]

	if found {
		repo.m[metric.Key] -= metric.Value
	} else {
		repo.m[metric.Key] = -metric.Value
	}
	repo.mutex.Unlock()
}

func (repo *InMemoryRepo) GetValue(key string) int {
	return repo.m[key]
}
