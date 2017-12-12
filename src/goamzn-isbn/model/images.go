package model

import "encoding/xml"

type ImageSets struct {
	XMLName   xml.Name   `xml:"ImageSets"`
	ImageSets []ImageSet `xml:"ImageSets"`
}

type ImageSet struct {
	XMLName        xml.Name       `xml:"ImageSet"`
	SmallImage     SmallImage     `xml:"SmallImage"`
	MediumImage    MediumImage    `xml:"MediumImage"`
	LargeImage     LargeImage     `xml:"LargeImage"`
	SwatchImage    SwatchImage    `xml:"SwatchImage"`
	ThumbnailImage ThumbnailImage `xml:"ThumbnailImage"`
	TinyImage      TinyImage      `xml:"TinyImage"`
}

type SmallImage struct {
	XMLName xml.Name `xml:"SmallImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}

type MediumImage struct {
	XMLName xml.Name `xml:"MediumImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}

type LargeImage struct {
	XMLName xml.Name `xml:"LargeImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}

type SwatchImage struct {
	XMLName xml.Name `xml:"SwatchImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}

type ThumbnailImage struct {
	XMLName xml.Name `xml:"ThumbnailImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}

type TinyImage struct {
	XMLName xml.Name `xml:"TinyImage"`
	URL     string   `xml:"URL"`
	Height  int64    `xml:"Height"`
	Width   int64    `xml:"Width"`
}
