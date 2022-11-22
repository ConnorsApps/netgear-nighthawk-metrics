package utils

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type FooCollector struct {
	fooMetric *prometheus.Desc
}

func Collector() *FooCollector {
	fooMetric := prometheus.NewDesc("foo_metric", "Some new metric", nil, nil)

	return &FooCollector{
		fooMetric,
	}
}

func (collector *FooCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.fooMetric
}

func (collector *FooCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue = 1

	log.Println("Value find")

	ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.CounterValue, float64(metricValue))

}
