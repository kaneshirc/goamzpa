// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func Test_NewRequest(t *testing.T) {
	file := "./testfiles/Test_NewRequest.xml"
	if isFileExists(file) {
		return
	}

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemLookup("ItemAttributes", "ASIN", "0141033576")

	checkAndWriteFile(t, file, response, err)
}

func Test_NewRequestWithClient(t *testing.T) {
	file := "./testfiles/Test_NewRequestWithClient.xml"
	if isFileExists(file) {
		return
	}

	request := NewRequestWithClient(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US", &http.Client{})
	response, err := request.ItemLookup("ItemAttributes", "ASIN", "0141033576")

	checkAndWriteFile(t, file, response, err)
}

func Test_ItemLookup(t *testing.T) {
	file := "./testfiles/Test_ItemLookup.xml"
	if isFileExists(file) {
		return
	}

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemLookup("ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews", "ASIN", "0141033576", "0615314465")

	checkAndWriteFile(t, file, response, err)
}

func Test_ItemSearch_All(t *testing.T) {
	file := "./testfiles/Test_ItemSearch_All.xml"
	if isFileExists(file) {
		return
	}

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemSearch("ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews", "All", "", "golang")

	checkAndWriteFile(t, file, response, err)
}

func Test_ItemSearch_Books(t *testing.T) {
	file := "./testfiles/Test_ItemSearch_Books.xml"
	if isFileExists(file) {
		return
	}

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemSearch("ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews", "Books", "salesrank", "golang")

	checkAndWriteFile(t, file, response, err)
}

func Test_Unmarshal_ItemLookup(t *testing.T) {
	content, err := ioutil.ReadFile("./testfiles/Test_ItemLookup.xml")
	if content == nil || err != nil {
		t.Error(err)
	}

	itemLookupResponse := ItemLookupResponse{}
	err = xml.Unmarshal(content, &itemLookupResponse)
	if err != nil {
		t.Error(err)
	}

	if !itemLookupResponse.Request.IsValid {
		t.Error("IsValid:", itemLookupResponse.Request.IsValid)
	}

	for i, item := range itemLookupResponse.Items {
		fmt.Println("-------------------")
		fmt.Printf("Item %d\n", i)
		fmt.Println("-------------------")
		fmt.Printf("ASIN: %s\n", item.ASIN)
		fmt.Printf("DetailPageURL: %s\n", item.DetailPageURL)
		fmt.Println()

		fmt.Println("  -------------------")
		fmt.Println("  ItemLinks")
		fmt.Println("  -------------------")
		for j, itemLink := range item.ItemLinks {
			fmt.Printf("  ItemLink %d\n", j)
			fmt.Printf("  Description: %s\n", itemLink.Description)
			fmt.Printf("  URL: %s\n", itemLink.URL)
			fmt.Println("  -----")
		}
		fmt.Println()

		fmt.Println("  -------------------")
		fmt.Println("  SmallImage")
		fmt.Println("  -------------------")
		printImage(&item.SmallImage)
		fmt.Println("  -------------------")
		fmt.Println("  MediumImage")
		fmt.Println("  -------------------")
		printImage(&item.MediumImage)
		fmt.Println("  -------------------")
		fmt.Println("  LargeImage")
		fmt.Println("  -------------------")
		printImage(&item.LargeImage)
		fmt.Println()

		itemAttributes := &item.ItemAttributes
		fmt.Println("  -------------------")
		fmt.Println("  ItemAttributes")
		fmt.Println("  -------------------")
		fmt.Printf("  EAN: %s\n", itemAttributes.EAN)
		fmt.Printf("  EANs: %s\n", strings.Join(itemAttributes.EANs, ","))
		fmt.Printf("  ISBN: %s\n", itemAttributes.ISBN)
		fmt.Printf("  UPC: %s\n", itemAttributes.UPC)
		fmt.Printf("  UPCs: %s\n", strings.Join(itemAttributes.UPCs, ","))
		fmt.Printf("  Title: %s\n", itemAttributes.Title)
		fmt.Printf("  Label: %s\n", itemAttributes.Label)
		fmt.Printf("  Author: %s\n", itemAttributes.Author)
		fmt.Printf("  Manufacturer: %s\n", itemAttributes.Manufacturer)
		fmt.Printf("  Publisher: %s\n", itemAttributes.Publisher)
		fmt.Printf("  Studio: %s\n", itemAttributes.Studio)
		fmt.Printf("  Brand: %s\n", itemAttributes.Brand)
		fmt.Printf("  ProductGroup: %s\n", itemAttributes.ProductGroup)
		fmt.Printf("  ProductTypeName: %s\n", itemAttributes.ProductTypeName)
		fmt.Printf("  Binding: %s\n", itemAttributes.Binding)
		fmt.Printf("  Edition: %s\n", itemAttributes.Edition)
		fmt.Printf("  PublicationDate: %s\n", itemAttributes.PublicationDate)
		fmt.Printf("  Feature: %s\n", itemAttributes.Feature)

		fmt.Println("    -------------------")
		fmt.Println("    Languages")
		fmt.Println("    -------------------")
		for j, language := range itemAttributes.Languages {
			fmt.Printf("    Language %d\n", j)
			fmt.Printf("    Name: %s\n", language.Name)
			fmt.Printf("    Type: %s\n", language.Type)
			fmt.Println("    -----")
		}

		fmt.Println("    -------------------")
		fmt.Println("    ItemDimensions")
		fmt.Println("    -------------------")
		fmt.Printf("    Height: %d\n", itemAttributes.ItemDimensions.Height)
		fmt.Printf("    Width: %d\n", itemAttributes.ItemDimensions.Width)
		fmt.Printf("    Length: %d\n", itemAttributes.ItemDimensions.Length)
		fmt.Printf("    Weight: %d\n", itemAttributes.ItemDimensions.Weight)

		fmt.Println("    -------------------")
		fmt.Println("    PackageDimensions")
		fmt.Println("    -------------------")
		fmt.Printf("    Height: %d\n", itemAttributes.PackageDimensions.Height)
		fmt.Printf("    Width: %d\n", itemAttributes.PackageDimensions.Width)
		fmt.Printf("    Length: %d\n", itemAttributes.PackageDimensions.Length)
		fmt.Printf("    Weight: %d\n", itemAttributes.PackageDimensions.Weight)
		fmt.Println()

		fmt.Println("  -------------------")
		fmt.Println("  CustomerReviews")
		fmt.Println("  -------------------")
		fmt.Printf("  CustomerReviewsUrl: %s\n", item.CustomerReviewsURL)
	}

}

func isFileExists(filename string) bool {
	os.Mkdir("testfiles", 0777)
	_, err := os.Stat(filename)
	return err == nil
}

func checkAndWriteFile(t *testing.T, file string, response []byte, err error) {
	if response == nil || err != nil {
		t.Error(err)
	} else {
		if err = ioutil.WriteFile(file, response, os.ModePerm); err != nil {
			t.Error(err)
		}
	}
}

func printImage(image *Image) {
	fmt.Printf("  URL: %s\n", image.URL)
	fmt.Printf("  Height: %d\n", image.Height)
	fmt.Printf("  Width: %d\n", image.Width)
}
