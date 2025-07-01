package adapter

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/prometheus/client_golang/prometheus"
)

// Target interface
type Metrics interface {
	IncCounter(name string, tags map[string]string)
	Observe(name string, value float64)
}

// Adaptee #1: Prometheus
type PromMetrics struct {
	counters   map[string]prometheus.Counter
	histograms map[string]prometheus.Histogram
}

func NewPromMetrics() *PromMetrics {
	return &PromMetrics{
		counters:   make(map[string]prometheus.Counter),
		histograms: make(map[string]prometheus.Histogram),
	}
}

func (m *PromMetrics) IncCounter(name string, tags map[string]string) {
	c, ok := m.counters[name]
	if !ok {
		c = prometheus.NewCounter(prometheus.CounterOpts{Name: name})
		prometheus.MustRegister(c)
		m.counters[name] = c
	}
	c.Inc()
}

func (m *PromMetrics) Observe(name string, value float64) {
	h, ok := m.histograms[name]
	if !ok {
		h = prometheus.NewHistogram(prometheus.HistogramOpts{Name: name})
		prometheus.MustRegister(h)
		m.histograms[name] = h
	}
	h.Observe(value)
}

// Adaptee #2: StatsD
type StatsdMetrics struct {
	client statsd.Statter
}

func NewStatsdMetrics(addr string) (*StatsdMetrics, error) {
	c, err := statsd.NewClient(addr, "")
	if err != nil {
		return nil, err
	}
	return &StatsdMetrics{client: c}, nil
}

func (s *StatsdMetrics) IncCounter(name string, tags map[string]string) {
	_ = s.client.Inc(name, 1, 1.0)
}

func (s *StatsdMetrics) Observe(name string, value float64) {
	_ = s.client.Timing(name, int64(value), 1.0)
}

// Client wiring
func RunMetricsAdapterDemo() {
	var m Metrics

	if useProm := false; useProm {
		m = NewPromMetrics()
	} else {
		m, _ = NewStatsdMetrics("127.0.0.1:8125")
	}

	m.IncCounter("requests_total", nil)
	m.Observe("request_latency_ms", 123.4)
}
