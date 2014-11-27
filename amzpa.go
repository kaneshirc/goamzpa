// Copyright 2012 Marco Dinacci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goamzpa provides functionality for using the
// Amazon Product Advertising service.

package goamzpa

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var service_domains = map[string]string{
	"CA": "ecs.amazonaws.ca",
	"CN": "webservices.amazon.cn",
	"DE": "ecs.amazonaws.de",
	"ES": "webservices.amazon.es",
	"FR": "ecs.amazonaws.fr",
	"IT": "webservices.amazon.it",
	"JP": "ecs.amazonaws.jp",
	"UK": "ecs.amazonaws.co.uk",
	"US": "ecs.amazonaws.com",
}

type AmazonRequest struct {
	AccessKeyID     string
	AccessKeySecret string
	AssociateTag    string
	Region          string
	Version         string
}

// Perform an ItemLookup request.
//
// Usage:
// response,err := request.ItemLookup("Medium,Accessories", "ASIN", "01289328", "2837423")
func (request *AmazonRequest) ItemLookup(responseGroups string, idType string, itemIds ...string) ([]byte, ItemLookupResponse, error) {
	arguments := make(map[string]string)
	arguments["Operation"] = "ItemLookup"
	arguments["IdType"] = idType
	arguments["ItemId"] = strings.Join(itemIds, ",")

	requestURL := request.buildURL(arguments, responseGroups)
	contents, err := doRequest(requestURL)

	response := ItemLookupResponse{}

	if err == nil {
		err = xml.Unmarshal(contents, &response)
	}

	return contents, response, err
}

// Perform an ItemLookup request.
//
// Usage:
// response,err := request.ItemSearch("Medium,Accessories", "All", "golang")
func (request *AmazonRequest) ItemSearch(responseGroups string, searchIndex string, keywords string) ([]byte, ItemSearchResponse, error) {
	arguments := make(map[string]string)
	arguments["Operation"] = "ItemSearch"
	arguments["SearchIndex"] = searchIndex
	arguments["Keywords"] = keywords

	requestURL := request.buildURL(arguments, responseGroups)
	contents, err := doRequest(requestURL)

	response := ItemSearchResponse{}

	if err == nil {
		err = xml.Unmarshal(contents, &response)
	}

	return contents, response, err
}

// Build and sign amazon specific URL
//
// Usage:
// url := request.buildURL(arguments, responseGroup)
func (request *AmazonRequest) buildURL(arguments map[string]string, responseGroups string) string {
	now := time.Now()
	arguments["AWSAccessKeyId"] = request.AccessKeyID
	arguments["Version"] = request.Version
	arguments["Timestamp"] = now.Format(time.RFC3339)
	arguments["Service"] = "AWSEcommerceService"
	arguments["AssociateTag"] = request.AssociateTag
	arguments["ResponseGroup"] = responseGroups

	// Sort the keys otherwise Amazon hash will be
	// different from mine and the request will fail
	keys := make([]string, 0, len(arguments))
	for argument := range arguments {
		keys = append(keys, argument)
	}
	sort.Strings(keys)

	// There's probably a more efficient way to concatenate strings, not a big deal though.
	var queryString string
	for _, key := range keys {
		escapedArg := url.QueryEscape(arguments[key])
		queryString += fmt.Sprintf("%s=%s", key, escapedArg)

		// Add '&' but only if it's not the the last argument
		if key != keys[len(keys)-1] {
			queryString += "&"
		}
	}

	// Hash & Sign
	domain := service_domains[request.Region]

	data := "GET\n" + domain + "\n/onca/xml\n" + queryString
	hash := hmac.New(sha256.New, []byte(request.AccessKeySecret))
	hash.Write([]byte(data))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
	queryString = fmt.Sprintf("%s&Signature=%s", queryString, signature)

	// build request URL
	requestURL := fmt.Sprintf("http://%s/onca/xml?%s", domain, queryString)
	return requestURL
}

// TODO add "Accept-Encoding": "gzip" and override UserAgent
// which is set to Go http package.
func doRequest(requestURL string) ([]byte, error) {
	var httpResponse *http.Response
	var err error
	var contents []byte

	httpResponse, err = http.Get(requestURL)

	if err != nil {
		return []byte(""), err
	}

	contents, err = ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()

	return contents, err
}
