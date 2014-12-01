// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"encoding/xml"
)

//+-------------------------------------------------
//| Request
//+-------------------------------------------------

type Request struct {
	XMLName xml.Name `xml:"Request"`
	IsValid bool     `xml:"IsValid"`
}

//+-------------------------------------------------
//| Response
//+-------------------------------------------------

// Mapping xml to struct.
//
// Usage:
//    import "encoding/xml"
//
//    itemLookupResponse := ItemLookupResponse{}
//    xml.Unmarshal(data, &itemLookupResponse)
type ItemLookupResponse struct {
	ItemResponseBase
	XMLName xml.Name `xml:"ItemLookupResponse"`
}

// Mapping xml to struct.
//
// Usage:
//    import "encoding/xml"
//
//    itemLookupResponse := ItemLookupResponse{}
//    xml.Unmarshal(data, &itemLookupResponse)
type ItemSearchResponse struct {
	ItemResponseBase
	XMLName              xml.Name `xml:"ItemSearchResponse"`
	TotalResults         uint16   `xml:"Items>TotalResults"`
	TotalPages           uint16   `xml:"Items>TotalPages"`
	MoreSearchResultsUrl string   `xml:"Items>MoreSearchResultsUrl"`
}

type ItemResponseBase struct {
	Items   []Item  `xml:"Items>Item"`
	Request Request `xml:"Items>Request"`
}

//+-------------------------------------------------
//| ResponseGroup
//+-------------------------------------------------

// ResponseGroup
type ItemAttributes struct {
	XMLName           xml.Name `xml:"ItemAttributes"`
	EAN               string
	EANs              []string `xml:"EANList>EANListElement"`
	ISBN              string
	UPC               string
	UPCs              []string `xml:"UPCList>UPCListElement"`
	Title             string
	Label             string
	Author            string
	Manufacturer      string
	Publisher         string
	Studio            string
	Brand             string
	ProductGroup      string
	ProductTypeName   string
	Binding           string
	Edition           string
	PublicationDate   string
	Feature           string
	Languages         []Language `xml:"Languages>Language"`
	ItemDimensions    Dimension  `xml:"ItemDimensions"`
	PackageDimensions Dimension  `xml:"PackageDimensions"`
}

//+-------------------------------------------------
//| Response Parts
//+-------------------------------------------------

type Dimension struct {
	Height uint16
	Width  uint16
	Length uint16
	Weight uint16
}
type Language struct {
	Name string
	Type string
}

type Image struct {
	URL    string
	Height uint16
	Width  uint16
}

type ItemLink struct {
	Description string
	URL         string
}

type Item struct {
	XMLName            xml.Name `xml:"Item"`
	ASIN               string
	URL                string
	DetailPageURL      string
	ItemAttributes     ItemAttributes `xml:"ItemAttributes"`
	SmallImage         Image          `xml:"SmallImage"`
	MediumImage        Image          `xml:"MediumImage"`
	LargeImage         Image          `xml:"LargeImage"`
	ItemLinks          []ItemLink     `xml:"ItemLinks>ItemLink"`
	CustomerReviewsURL string         `xml:"CustomerReviews>IFrameURL"`
}
