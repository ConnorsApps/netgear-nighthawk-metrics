package main

import (
	"log"

	"github.com/ConnorsApps/netgear-nighthawk-metrics/utils"
)

func main() {
	login := utils.ParseArgs()

	metrics := utils.MetricsRequest(login)

	log.Println("resp", *metrics)
}
