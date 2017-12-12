package errors

import "encoding/xml"

type ItemLookupErrorResponse struct {
	XMLName   xml.Name `xml:"ItemLookupErrorResponse"`
	Error     Error    `xml:"Error"`
	RequestID string   `xml:"RequestID"`
}

type Error struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}
