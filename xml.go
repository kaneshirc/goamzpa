// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"encoding/xml"
)

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
	XMLName       xml.Name `xml:"Item"`
	ASIN          string
	URL           string
	DetailPageURL string
	Title         string     `xml:"ItemAttributes>Title"`
	Author        string     `xml:"ItemAttributes>Author"`
	Price         string     `xml:"ItemAttributes>ListPrice>FormattedPrice"`
	PriceRaw      string     `xml:"ItemAttributes>ListPrice>Amount"`
	SmallImage    Image      `xml:"SmallImage"`
	MediumImage   Image      `xml:"MediumImage"`
	LargeImage    Image      `xml:"LargeImage"`
	ItemLinks     []ItemLink `xml:"ItemLinks>ItemLink"`
}

type Request struct {
	XMLName           xml.Name          `xml:"Request"`
	IsValid           bool              `xml:"IsValid"`
	ItemLookupRequest ItemLookupRequest `xml:"ItemLookupRequest"`
}

type ItemLookupRequest struct {
	XMLName        xml.Name `xml:"ItemLookupRequest"`
	IdType         string
	ItemIds        []string `xml:"ItemId"`
	ResponseGroups []string `xml:"ResponseGroup"`
	VariationPage  string
}

type ItemResponseBase struct {
	Items   []Item  `xml:"Items>Item"`
	Request Request `xml:"Items>Request"`
}

type ItemLookupResponse struct {
	ItemResponseBase
	XMLName xml.Name `xml:"ItemLookupResponse"`
}

type ItemSearchResponse struct {
	ItemResponseBase
	XMLName              xml.Name `xml:"ItemSearchResponse"`
	TotalResults         uint16   `xml:"Items>TotalResults"`
	TotalPages           uint16   `xml:"Items>TotalPages"`
	MoreSearchResultsUrl string   `xml:"Items>MoreSearchResultsUrl"`
}
