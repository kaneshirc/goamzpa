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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var hosts = map[string]string{
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
	Service         string
	Version         string
	Client          *http.Client
}

// Create a new AmazonRequest initialized with the given parameters
func NewRequest(accessKeyID string, accessKeySecret string, associateTag string, region string) *AmazonRequest {
	return &AmazonRequest{accessKeyID, accessKeySecret, associateTag, region, "AWSEcommerceService", "2013-08-01", &http.Client{}}
}

// Create a new AmazonRequest initialized with the given parameters
func NewRequestWithClient(accessKeyID string, accessKeySecret string, associateTag string, region string, client *http.Client) *AmazonRequest {
	return &AmazonRequest{accessKeyID, accessKeySecret, associateTag, region, "AWSEcommerceService", "2013-08-01", client}
}

// Perform an ItemLookup request.
//
// Usage:
//    response,err := request.ItemLookup("Medium,Accessories", "ASIN", "01289328", "2837423")
func (request *AmazonRequest) ItemLookup(responseGroups string, idType string, itemIds ...string) ([]byte, error) {
	params := make(map[string]string)
	params["Operation"] = "ItemLookup"
	params["IdType"] = idType
	params["ItemId"] = strings.Join(itemIds, ",")

	requestURL := request.buildURL(params, responseGroups)
	return request.doRequest(requestURL)
}

// Perform an ItemLookup request.
//
// Usage:
//    response,err := request.ItemSearch("Medium,Accessories", "All", "golang")
func (request *AmazonRequest) ItemSearch(responseGroups string, searchIndex string, keywords string) ([]byte, error) {
	params := make(map[string]string)
	params["Operation"] = "ItemSearch"
	params["SearchIndex"] = searchIndex
	params["Keywords"] = keywords

	requestURL := request.buildURL(params, responseGroups)
	return request.doRequest(requestURL)
}

// Build and sign amazon specific URL
//
// Usage:
//    url := request.buildURL(params, responseGroup)
func (request *AmazonRequest) buildURL(params map[string]string, responseGroups string) string {
	now := time.Now()
	params["AWSAccessKeyId"] = request.AccessKeyID
	params["Version"] = request.Version
	params["Timestamp"] = now.Format(time.RFC3339)
	params["Service"] = request.Service
	params["AssociateTag"] = request.AssociateTag
	params["ResponseGroup"] = responseGroups

	// Sort the keys otherwise Amazon hash will be different from mine and the request will fail
	keys := make([]string, 0, len(params))
	for argument := range params {
		keys = append(keys, argument)
	}
	sort.Strings(keys)

	// TODO There's probably a more efficient way to concatenate strings, not a big deal though.
	//      もっと効率のいい文字列連結方法があるはず
	var queryString string
	for _, key := range keys {
		escapedArg := url.QueryEscape(params[key])
		queryString += fmt.Sprintf("%s=%s", key, escapedArg)

		// Add '&' but only if it's not the the last argument
		if key != keys[len(keys)-1] {
			queryString += "&"
		}
	}

	// Hash & Sign
	host := hosts[request.Region]
	path := "/onca/xml"
	data := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, queryString)
	hash := hmac.New(sha256.New, []byte(request.AccessKeySecret))
	hash.Write([]byte(data))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
	signedQueryString := fmt.Sprintf("%s&Signature=%s", queryString, signature)

	// build request URL
	requestURL := fmt.Sprintf("http://%s/onca/xml?%s", host, signedQueryString)
	return requestURL
}

// which is set to Go http package.
func (request *AmazonRequest) doRequest(requestURL string) ([]byte, error) {

	httpResponse, err := request.Client.Get(requestURL)
	defer httpResponse.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(httpResponse.Body)
}
