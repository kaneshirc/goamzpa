// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"io/ioutil"
	"net/http"
	"os"
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
