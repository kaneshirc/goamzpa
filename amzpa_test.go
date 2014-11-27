// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"fmt"
	"os"
	"testing"
)

func TestItemLookup(t *testing.T) {
	// Complete these variables with your credentials
	accessKey := os.Getenv("ACCESS_KEY")
	accessSecret := os.Getenv("ACCESS_SECRET")
	associateTag := os.Getenv("ASSOCIATE_TAG")
	region := "UK"
	version := "2013-08-01"

	request := &AmazonRequest{accessKey, accessSecret, associateTag, region, version}

	asins := []string{"0141033576,0615314465,1470057719"}

	responseGroups := "ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews"
	itemsType := "ASIN"
	contents, response, err := request.ItemLookup(responseGroups, itemsType, asins...)

	if err == nil && response.Request.IsValid {
		fmt.Printf("\n%s\n", string(contents))

		fmt.Println("-------------------")
		fmt.Println("Request information")
		fmt.Println("-------------------")
		fmt.Printf("IdType: %s\n", response.Request.ItemLookupRequest.IdType)
		fmt.Printf("ItemIds: %s\n", response.Request.ItemLookupRequest.ItemIds)
		fmt.Printf("ResponseGroups: %s\n", response.Request.ItemLookupRequest.ResponseGroups)
		fmt.Printf("VariationPage: %s\n", response.Request.ItemLookupRequest.VariationPage)
		fmt.Println("-------------------")

		for count, item := range response.Items {
			fmt.Println("-------------------")
			fmt.Printf("Item %d\n", count+1)
			fmt.Println("-------------------")
			fmt.Printf("ASIN: %s\n", item.ASIN)
			fmt.Printf("Title: %s\n", item.Title)
			fmt.Printf("DetailPageURL: %s\n", item.DetailPageURL)
			fmt.Printf("Author: %s\n", item.Author)
			fmt.Printf("Price: %s\n", item.Price)
			fmt.Printf("Small Image URL: %s\n", item.SmallImage.URL)
			fmt.Printf("Medium Image URL: %s\n", item.MediumImage.URL)
			fmt.Printf("Large Image URL: %s\n", item.LargeImage.URL)
		}
	} else {
		t.Error(err)
	}
}

func TestItemSearch(t *testing.T) {
	// Complete these variables with your credentials
	accessKey := os.Getenv("ACCESS_KEY")
	accessSecret := os.Getenv("ACCESS_SECRET")
	associateTag := os.Getenv("ASSOCIATE_TAG")
	region := "US"
	version := "2013-08-01"

	request := &AmazonRequest{accessKey, accessSecret, associateTag, region, version}
	responseGroups := "ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews"

	contents, response, err := request.ItemSearch(responseGroups, "All", "golang")

	if err == nil && response.Request.IsValid {
		fmt.Printf("\n%s\n", string(contents))

		fmt.Println("-------------------")
		fmt.Println("Request information")
		fmt.Println("-------------------")
		for count, item := range response.Items {
			fmt.Println("-------------------")
			fmt.Printf("Item %d\n", count+1)
			fmt.Println("-------------------")
			fmt.Printf("ASIN: %s\n", item.ASIN)
			fmt.Printf("Title: %s\n", item.Title)
		}

	} else {
		t.Error(err)
	}
}
