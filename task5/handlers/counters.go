package counters

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	Calls prometheus.Counter
}

func New(namespace string, name string) *Metric {
	return &Metric{Calls: promauto.NewCounter(prometheus.CounterOpts{Namespace: namespace, Name: name})}
}
