package readhtml

import (
	"bytes"
	"io"
	"log"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

func ReadHotelSectionListAndFacility(body io.Reader) (string, string) {

	htmltc, xmlerr := html.Parse(body)
	if xmlerr != nil {
		log.Fatal(xmlerr)
	}
	q := "n/a"
	detail := "n/a"
	xpath := `//*[@id="HotelSectionList"]`
	node := htmlquery.FindOne(htmltc, xpath)

	section := htmlquery.InnerText(node)
	hotel_details := `//*[@id="hotel-details"]/div[1]/ul/li`

	details := bytes.Buffer{}
	htmlquery.FindEach(htmltc, hotel_details, func(i int, node *html.Node) {

		lable := htmlquery.FindOne(node, "//label")

		details.WriteString(htmlquery.InnerText(lable))
		details.WriteString(":")
		span := "//span"
		htmlquery.FindEach(node, span, func(j int, spanNode *html.Node) {

			s := htmlquery.FindOne(spanNode, "//span")
			details.WriteString(htmlquery.InnerText(s))
			details.WriteString(",")
		})
	})
	detail = details.String()
	if section != "" {
		q = section
	}
	return q, detail

}
