package main

import (
	"log"
	"net/http"

	"github.com/ConnorsApps/netgear-nighthawk-metrics/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	args := utils.ParseArgs()

	foo := utils.PortsCollector(args)
	prometheus.MustRegister(foo)

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":"+args.Port, nil))
}
