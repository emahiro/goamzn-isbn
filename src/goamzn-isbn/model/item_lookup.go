package model

import "encoding/xml"

type ItemLookupResponse struct {
	XMLName xml.Name `xml:"ItemLookupResponse"`
	Items   Items    `xml:"Items"`
}

type Items struct {
	XMLName xml.Name `xml:"Items"`
	Request Request  `xml:"Request"`
	Item    []Item   `xml:"Item"`
}

type Request struct {
	XMLName           xml.Name          `xml:"Request"`
	IsValid           bool              `xml:"IsValid"`
	ItemLookupRequest ItemLookupRequest `xml:"ItemLookupRequest"`
}

type ItemLookupRequest struct {
	XMLName        xml.Name `xml:"ItemLookupRequest"`
	IdType         string   `xml:"IdType"`
	ItemId         int64    `xml:"ItemId"`
	ResponseGroup  string   `xml:"ResponseGroup"`
	SearchIndex    string   `xml:"SearchIndex"`
	VariactionPage string   `xml:"VariationPage"`
}

type Item struct {
	XMLName        xml.Name       `xml:"Item"`
	ASIN           int64          `xml:"ASIN"`
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

type ItemAttributes struct {
	XMLName      xml.Name `xml:"ItemAttributes"`
	ISBN         string   `xml:"ISBN"`
	Title        string   `xml:"Title"`
	Author       []string `xml:"Author"`
	Manufacturer string   `xml:"Manufacturer"`
	ProductGroup string   `xml:"ProductGroup"`
}
