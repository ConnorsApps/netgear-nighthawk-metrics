package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ConnorsApps/netgear-nighthawk-metrics/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const METRICS_ENDPOINT = "/metrics"
const VERSION = "1.0.0"

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	args := utils.ParseArgs()

	foo := utils.PortsCollector(args)
	prometheus.MustRegister(foo)

	fmt.Println("Listening at localhost:"+args.Port+METRICS_ENDPOINT, "version:", VERSION)

	http.Handle(METRICS_ENDPOINT, promhttp.Handler())
	http.Handle("/health", http.HandlerFunc(health))

	log.Panicln(http.ListenAndServe(":"+args.Port, nil))
}
