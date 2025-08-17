package main

import (
	"fmt"
	"net/url"
)

// Extracting different parts of the url-string
// [scheme://][userIfo@]host[:port][/path][?query][#fragment]

func main() {
	rawURL:="https://example.com:8080/path?query=param#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err!=nil{
		fmt.Println("🔴ERROR parsing URL:",err)
		return
	}
	fmt.Println("Full Parsed-URL:",parsedURL)
	fmt.Println("Host:",parsedURL.Host)
	fmt.Println("Port:",parsedURL.Port())
	fmt.Println("Path:",parsedURL.Path)
	fmt.Println("RawQuery:",parsedURL.RawQuery)
	fmt.Println("Fragment:",parsedURL.Fragment)

	// QUERY-PARAMS
	rawURL1:="https://example.com/path?name=John&age=30"
	parsedURL1,err:= url.Parse(rawURL1)
	if err!=nil{
		fmt.Println("🔴ERROR parsing URL:",err)
		return
	}

	queryParams:= parsedURL1.Query()
	fmt.Println("Query-Params:",queryParams) // map[age:[30] name:[John]]
	fmt.Println("Name:",queryParams.Get("name"))
	fmt.Println("Age:",queryParams.Get("age"))

	// building URL
	baseURL:= &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}

	query:= baseURL.Query()
	query.Set("name","Skyy")
	query.Set("age","29")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL:",baseURL.String())

	// Alternate WAY
	values:=url.Values{}

	// Add key-value pairs to the values obj{}
	values.Add("country","Germany")
	values.Add("capital","Berlin")
	values.Add("chancellor","Merz")
	values.Add("stdcode","49")

	// Encode
	encodedQuery:= values.Encode()
	fmt.Println(encodedQuery)

	// Build URL
	baseURL1:= "https://example.com/search"
	fullURL:= baseURL1+"?"+encodedQuery
	fmt.Println("Full URL:",fullURL)

}