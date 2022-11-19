package main

import (
	"github.com/ConnorsApps/netgear-nighthawk-metrics/utils"
)

func main() {
	login := utils.ParseArgs()

	response := utils.MetricsRequest(login)

	utils.ResponseParser(response)
}
