package entity

import "time"

type Metric struct {
	Key       string
	Value     int
	CreatedAt time.Time
}
