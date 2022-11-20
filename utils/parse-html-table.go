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
	Status                    string
	TransmittedPackets        int
	ReceivedPackets           int
	Collisions                int
	TransmittedBytesPerSecond int
	ReceivedBytesPerSecond    int
	Uptime                    string
}

func toInt(str string) int {
	if str == "--" {
		return 0
	} else {
		intVar, err := strconv.Atoi(str)
		if err != nil {
			log.Panicln("Unable to convert", str, "to int", err)
		}
		return intVar
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
			stat.Status = text
		} else if y == 2 {
			stat.TransmittedPackets = toInt(text)
		} else if y == 3 {
			stat.ReceivedPackets = toInt(text)
		} else if y == 4 {
			stat.Collisions = toInt(text)
		} else if y == 5 {
			stat.TransmittedBytesPerSecond = toInt(text)
		} else if y == 6 {
			stat.ReceivedBytesPerSecond = toInt(text)
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
