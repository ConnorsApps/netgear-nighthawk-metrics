package utils

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseDoc(response io.ReadCloser) (string, *goquery.Document) {
	defer response.Close()

	bodyBytes, err := io.ReadAll(response)
	if err != nil {
		log.Fatalln(err)
	}
	bodyString := string(bodyBytes)
	if bodyString == "" {
		log.Fatalln("No body found")
	}
	parsedHtmlReader := strings.NewReader(bodyString)

	doc, err := goquery.NewDocumentFromReader(parsedHtmlReader)
	if err != nil {
		log.Fatalln(err)
	}
	return bodyString, doc
}
func systemUptime(body string) string {
	// <!> 14 days 02:24:56<!>

	afterFirstPoint := strings.Split(body, "<!>")[1]
	beforeNextPoint := strings.Split(afterFirstPoint, "<!>")[0]
	systemUptime := strings.TrimSpace(beforeNextPoint)

	return systemUptime
}
func ResponseParser(response io.ReadCloser) {
	body, doc := parseDoc(response)

	routerTitle := doc.Find("title").First().Text()

	uptime := systemUptime(body)

	log.Println("routerTitle", routerTitle, "upTime", uptime)
}
