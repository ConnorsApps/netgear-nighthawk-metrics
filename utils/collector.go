package utils

import (
	"github.com/prometheus/client_golang/prometheus"
)

var appArgs AppArgs

type RouterCollector struct {
	fooMetric *prometheus.GaugeVec
}

func (collector *RouterCollector) Describe(ch chan<- *prometheus.Desc) {
	collector.fooMetric.Describe(ch)
}

func (collector *RouterCollector) Collect(ch chan<- prometheus.Metric) {
	// var metricValue = 1

	response := RouterRequest(appArgs)

	stats := PraseHtml(response)

	// collector.fooMetric.withlabels
	// stats.Ports
	// collector.fooMetric, prometheus.CounterValue, float64(metricValue)

	// type PortStats struct {
	// 	Port                      string
	// 	ThroughputStatus          int
	// 	Status                    string
	// 	TransmittedPackets        int
	// 	ReceivedPackets           int
	// 	Collisions                int
	// 	TransmittedBytesPerSecond int
	// 	ReceivedBytesPerSecond    int
	// 	Uptime                    string
	// }
	for _, port := range stats.Ports {
		// labels := [2]string{port.Port, port.Status}

		collector.fooMetric.WithLabelValues(port.Port, port.Status).Set(float64(1))

	}
	collector.fooMetric.Collect(ch)

}

func PortsCollector(args AppArgs) *RouterCollector {
	appArgs = args
	fooMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "mynamespace",
		Subsystem: "client",
		Name:      "info",
		Help:      "something",
	}, []string{"port", "status"})

	return &RouterCollector{
		fooMetric,
	}
}
