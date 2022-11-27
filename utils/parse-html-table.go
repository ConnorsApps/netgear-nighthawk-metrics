package utils

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const OLD_TABLE_SELECTOR = "body > table:nth-child(2) > tbody > tr:nth-child(3) > td > table"
const NEW_TABLE_SELECTOR = "body > div > table"

type PortStats struct {
	Port                      string
	ThroughputStatus          float64
	Status                    string
	TransmittedPackets        float64
	ReceivedPackets           float64
	Collisions                float64
	TransmittedBytesPerSecond float64
	ReceivedBytesPerSecond    float64
	Uptime                    string
}

func toFloat(str string) float64 {
	if str == "--" {
		return 0
	} else {
		floatVar, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Panicln("Unable to convert", str, "to int", err)
		}
		return floatVar
	}
}
func ThroughputStatus(str string) float64 {
	// Example inputs 1000M/Full, Link Down, 750M, 1652M

	if strings.ToLower(str) == "link down" {
		return 0
	}
	str = strings.Replace(str, "/Full", "", 1)

	if strings.Contains(str, "M") {
		str = strings.Replace(str, "M", "", 1)

		return toFloat(str)
	} else {
		log.Panicln("Unknown Router Status", str)
		return 0
	}
}

func parseRow(col *goquery.Selection) PortStats {
	var stat PortStats

	// headers
	// Port Status TxPkts RxPkts Collisions Tx B/s Rx B/s Up Time

	col.Find("td").Each(func(y int, space *goquery.Selection) {

		text := strings.TrimSpace(space.Text())

		if y == 0 {
			stat.Port = text
		} else if y == 1 {
			stat.ThroughputStatus = ThroughputStatus(text)
			stat.Status = text
		} else if y == 2 {
			stat.TransmittedPackets = toFloat(text)
		} else if y == 3 {
			stat.ReceivedPackets = toFloat(text)
		} else if y == 4 {
			stat.Collisions = toFloat(text)
		} else if y == 5 {
			stat.TransmittedBytesPerSecond = toFloat(text)
		} else if y == 6 {
			stat.ReceivedBytesPerSecond = toFloat(text)
		} else if y == 7 {
			if text == "--" {
				stat.Uptime = ""
			} else {
				stat.Uptime = text
			}
		}
	})

	return stat
}

func ParseHtmlTable(doc *goquery.Document, isNewUI bool) []PortStats {

	var table *goquery.Selection

	if isNewUI {
		table = doc.Find(NEW_TABLE_SELECTOR)
	} else {
		table = doc.Find(OLD_TABLE_SELECTOR)
	}

	rows := table.Find("tr")

	var stats []PortStats

	rows.Each(func(i int, col *goquery.Selection) {
		if i != 0 {
			row := parseRow(col)
			stats = append(stats, row)
		}
	})

	return stats
}
