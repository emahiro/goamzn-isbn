package model

import "encoding/xml"

type ItemAttributes struct {
	XMLName           xml.Name          `xml:"ItemAttributes"`
	Author            []string          `xml:"Author"`
	Binding           string            `xml:"Binding"`
	EAN               int64             `xml:"EAN"`
	EANList           int64             `xml:"EANList>EANListElement"`
	IsAdultProduct    bool              `xml:"IsAdultProduct"`
	ISBN              string            `xml:"ISBN"`
	Label             string            `xml:"Label"`
	Languages         []Language        `xml:"Languages>Language"`
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
