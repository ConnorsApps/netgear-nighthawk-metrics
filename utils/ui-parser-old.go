package utils

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const OLD_TABLE_SELECTOR = "body > table:nth-child(2) > tbody > tr:nth-child(3) > td > table"

func UIParserOld(doc *goquery.Document) {

	table := doc.Find(OLD_TABLE_SELECTOR)

	rows := table.Find("tr")

	var headers []string

	rows.Each(func(i int, col *goquery.Selection) {
		// headers
		text := strings.TrimSpace(col.Text())
		log.Println("text'", text, "'")
		if i == 0 {
			headers = append(headers, text)
		}
	})

	log.Println("headers", headers)

	// rows.Map(func(i int, col *goquery.Selection) string {

	// 	if i == 0 {
	// 		col.

	// 	}

	// })

	// log.Println(text)
	//

}
