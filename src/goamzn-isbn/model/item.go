package model

import "encoding/xml"

type Item struct {
	XMLName        xml.Name       `xml:"Item"`
	ASIN           string         `xml:"ASIN"`
	DetailPageURL  string         `xml:"DetailPageURL"`
	ItemLinks      ItemLinks      `xml:"ItemLinks"`
	ItemAttributes ItemAttributes `xml:"ItemAttributes"`
}

type ItemLinks struct {
	XMLName  xml.Name   `xml:"ItemLinks"`
	ItemLink []ItemLink `xml:"ItemLink"`
}

type ItemLink struct {
	XMLName     xml.Name `xml:"ItemLink"`
	Description string   `xml:"Description"`
	URL         string   `xml:"URL"`
}
