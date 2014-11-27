# goamzpa

A BSD licensed [Go lang](http://golang.org) library to use the _Amazon Product API_. 
Also my first `Go` project.

At the moment it supports only `ItemLookup`. Everything can change, and
probably will, use at your own peril.

## Usage
    package main 
	import (
	    "fmt"
	    "github.com/mdinacci/goamzpa/amzpa"
	)

	func main() {
	    // Complete these variables with your credentials
	    accessKey := "ACCESS_KEY"
	    accessSecret := "ACCESS_SECRET"
	    associateTag := "ASSOCIATE_TAG"
	    region := "UK"

	    request := amzpa.NewRequest(accessKey, accessSecret , associateTag, region)
	    asins:= []string{"0141033576,0615314465,1470057719"}
	    
	    responseGroups := "Medium,Accessories"
	    itemsType := "ASIN"
	    response,err := request.ItemLookup(asins, responseGroups, itemsType)
	    
	    if err == nil && response.Request.IsValid {
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
				fmt.Printf("DetailPageURL: %s\n", item.DetailPageURL)
				fmt.Printf("Author: %s\n", item.Author)
				fmt.Printf("Price: %s\n", item.Price)
				fmt.Printf("Small Image URL: %s\n", item.SmallImage.URL)
				fmt.Printf("Medium Image URL: %s\n", item.MediumImage.URL)
				fmt.Printf("Large Image URL: %s\n", item.LargeImage.URL)
			}
		} else {
			fmt.Println(err)
		}
	}

 

## TODO
* [IN PROGRESS] Map the XML to a struct, so that the response is not just a big string
* Support more than one ResponseGroup
* Gzip compression

