package model

import (
	"encoding/xml"
)

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
	ItemId         string   `xml:"ItemId"`
	ResponseGroup  string   `xml:"ResponseGroup"`
	SearchIndex    string   `xml:"SearchIndex"`
	VariactionPage string   `xml:"VariationPage"`
}
