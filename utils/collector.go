package utils

import (
	"github.com/prometheus/client_golang/prometheus"
)

var appArgs AppArgs

type RouterCollector struct {
	port                      *prometheus.GaugeVec
	throughput                *prometheus.GaugeVec
	transmittedPackets        *prometheus.GaugeVec
	receivedPackets           *prometheus.GaugeVec
	collisions                *prometheus.GaugeVec
	transmittedBytesPerSecond *prometheus.GaugeVec
	receivedBytesPerSecond    *prometheus.GaugeVec
}

func (collector *RouterCollector) Describe(ch chan<- *prometheus.Desc) {
	collector.port.Describe(ch)
	collector.throughput.Describe(ch)
	collector.transmittedPackets.Describe(ch)
	collector.receivedPackets.Describe(ch)
	collector.collisions.Describe(ch)
	collector.transmittedBytesPerSecond.Describe(ch)
	collector.receivedBytesPerSecond.Describe(ch)
}

func (collector *RouterCollector) Collect(ch chan<- prometheus.Metric) {
	response := RouterRequest(appArgs)

	stats := PraseHtml(response)

	for _, port := range stats.Ports {
		collector.port.WithLabelValues(port.Port, port.Status).Set(float64(1))
		collector.throughput.WithLabelValues(port.Port).Set(port.ThroughputStatus)
		collector.transmittedPackets.WithLabelValues(port.Port).Set(port.TransmittedPackets)
		collector.receivedPackets.WithLabelValues(port.Port).Set(port.ReceivedPackets)
		collector.collisions.WithLabelValues(port.Port).Set(port.Collisions)
		collector.transmittedBytesPerSecond.WithLabelValues(port.Port).Set(port.TransmittedBytesPerSecond)
		collector.receivedBytesPerSecond.WithLabelValues(port.Port).Set(port.ReceivedBytesPerSecond)
	}

	collector.port.Collect(ch)
	collector.throughput.Collect(ch)
	collector.transmittedPackets.Collect(ch)
	collector.receivedPackets.Collect(ch)
	collector.collisions.Collect(ch)
	collector.transmittedBytesPerSecond.Collect(ch)
	collector.receivedBytesPerSecond.Collect(ch)
}

func PortsCollector(args AppArgs) *RouterCollector {
	appArgs = args
	namespace := "netgear"
	port := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "port",
		Help:      "Netgear Port status",
	}, []string{"port", "status"})

	throughput := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "throughput_status",
		Help:      "Netgear Port Throughput from port status in Megabytes",
	}, []string{"port"})

	transmitted := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "transmitted",
		Help:      "Number of trasmitted packets",
	}, []string{"port"})

	receivedPackets := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "received_packets",
		Help:      "Number of received packets",
	}, []string{"port"})

	collisions := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "collisions",
		Help:      "Number of collisions",
	}, []string{"port"})

	transmittedBytesPerSecond := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "transmitted_bytes_per_second",
		Help:      "Trasmitted Bytes Per Second",
	}, []string{"port"})

	receivedBytesPerSecond := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "received_bytes_per_second",
		Help:      "Received Bytes Per Second",
	}, []string{"port"})

	return &RouterCollector{
		port,
		throughput,
		transmitted,
		receivedPackets,
		collisions,
		transmittedBytesPerSecond,
		receivedBytesPerSecond,
	}
}
