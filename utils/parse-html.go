package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func systemUptime(body string) string {
	// <!> 14 days 02:24:56<!>

	afterFirstPoint := strings.Split(body, "<!>")[1]
	beforeNextPoint := strings.Split(afterFirstPoint, "<!>")[0]
	systemUptime := strings.TrimSpace(beforeNextPoint)

	return systemUptime
}

type Stats struct {
	RouterTitle string
	Ports       []PortStats
	Uptime      string
}

func PraseHtml(body string, doc *goquery.Document) Stats {
	routerTitle := doc.Find("title").First().Text()

	isNewUI := doc.Find(".table_header").Length() > 0

	portStats := ParseHtmlTable(doc, isNewUI)

	uptime := systemUptime(body)

	return Stats{
		RouterTitle: routerTitle,
		Ports:       portStats,
		Uptime:      uptime,
	}
}
