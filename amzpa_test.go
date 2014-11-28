// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_NewRequest(t *testing.T) {

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemLookup("ItemAttributes", "ASIN", "0141033576")

	if response == nil || err != nil {
		t.Error(err)
	}
}

func Test_NewRequestWithClient(t *testing.T) {

	request := NewRequestWithClient(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US", &http.Client{})
	response, err := request.ItemLookup("ItemAttributes", "ASIN", "0141033576")

	if response == nil || err != nil {
		t.Error(err)
	}
}

func Test_ItemLookup(t *testing.T) {

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemLookup("ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews", "ASIN", "0141033576", "0615314465")

	fmt.Println(string(response))
	if response == nil || err != nil {
		t.Error(err)
	}
}

func Test_ItemSearch(t *testing.T) {

	request := NewRequest(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), os.Getenv("ASSOCIATE_TAG"), "US")
	response, err := request.ItemSearch("ItemAttributes,Images,BrowseNodes,EditorialReview,Reviews", "All", "golang")

	fmt.Println(string(response))
	if err != nil {
		t.Error(err)
	}

}
