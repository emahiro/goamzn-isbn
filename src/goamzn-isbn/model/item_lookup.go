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

type ItemAttributes struct {
	XMLName           xml.Name          `xml:"ItemAttributes"`
	Author            []string          `xml:"Author"`
	Binding           string            `xml:"Binding"`
	EAN               int64             `xml:"EAN"`
	EANList           EANList           `xml:"EANList"`
	IsAdultProduct    bool              `xml:"IsAdultProduct"`
	ISBN              string            `xml:"ISBN"`
	Label             string            `xml:"Label"`
	Languages         Languages         `xml:"Languages"`
	Manufacturer      string            `xml:"Manufacturer"`
	NumberOfPages     int64             `xml:"NumberOfPages"`
	PackageDimensions PackageDimensions `xml:"PackageDimensions"`
	ProductGroup      string            `xml:"ProductGroup"`
	ProductTypeName   string            `xml:"ProductTypeName"`
	PublicationDate   string            `xml:"PublicationDate"`
	Publisher         string            `xml:"Publisher"`
	Studio            string            `xml:"Studio"`
	Title             string            `xml:"Title"`
}

type EANList struct {
	XMLName        xml.Name `xml:"EANList"`
	EANListElement int64    `xml:"EANListElement"`
}

type Languages struct {
	XMLName  xml.Name   `xml:"Languages"`
	Language []Language `xml:"Language"`
}

type Language struct {
	XMLName xml.Name `xml:"Language"`
	Name    string   `xml:"Name"`
	Type    string   `xml:"Type"`
}

type PackageDimensions struct {
	XMLName xml.Name `xml:"PackageDimensions"`
	Height  int64    `xml:"Height"`
	Length  int64    `xml:"Length"`
	Weight  int64    `xml:"Weight"`
	Width   int64    `xml:"Width"`
}
