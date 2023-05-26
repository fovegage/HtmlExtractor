package test

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"testing"
)

func ReadHtml() {
	sourceByte, _ := ioutil.ReadFile("/Users/gaozhe/GolandProjects/vpnbook/spider/HtmlExtractor/test/html/asos.html")
	dom, _ := goquery.NewDocumentFromReader(bytes.NewReader(sourceByte))
	coordinate, _ := dom.Find(".image-thumbnail").First().Attr("coordinate")
	println(coordinate)

}

func TestExtractor(t *testing.T) {
	ReadHtml()
}
