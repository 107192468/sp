package proxyurl

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v1"
)

func TestP(t *testing.T) {
	brokenHtml := `<!DOCTYPE html><html><body><h1 id="someid">My First Heading</h1><p>paragraph</body></html>`

	reader := strings.NewReader(brokenHtml)
	root, err := html.Parse(reader)

	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	html.Render(&b, root)
	fixedHtml := b.String()

	reader = strings.NewReader(fixedHtml)
	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	var xpath string
	xpath = `//h1[@id='someid']`
	path := xmlpath.MustCompile(xpath)
	if value, ok := path.String(xmlroot); ok {
		log.Println("Found:", value)
	}
}
