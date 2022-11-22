package main

import (
	"log"
	"net/http"

	"github.com/ConnorsApps/netgear-nighthawk-metrics/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const METRICS_ENDPOINT = "/metrics"

func main() {
	args := utils.ParseArgs()

	foo := utils.PortsCollector(args)
	prometheus.MustRegister(foo)

	log.Println("Listening at localhost:" + args.Port + METRICS_ENDPOINT)

	http.Handle(METRICS_ENDPOINT, promhttp.Handler())

	log.Fatal(http.ListenAndServe(":"+args.Port, nil))
}
